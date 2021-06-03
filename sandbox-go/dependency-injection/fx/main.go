package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/butuzov/sandbox/dependency-injection/fx/httphandler"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func main() {
	os.Exit(run())
}

func run() int {
	fx.New(
		fx.Provide(ProvideConfig),
		fx.Provide(ProvideLogger),
		fx.Provide(http.NewServeMux),
		fx.Provide(httphandler.New),
		// invokers
		fx.Invoke(registerHooks),
	).Run()

	return 0
}

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	cfg *Config,
	mux *httphandler.Handler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(cfg.ApplicationConfig.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return slogger
}

func ProvideConfig() *Config {
	var buf bytes.Buffer
	cnf := &Config{}

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Printf("config: %s", err)
		return nil
	}

	defer f.Close()

	if _, err := io.Copy(&buf, f); err != nil {
		log.Printf("reading: %s", err)
		return nil
	}

	if err := yaml.Unmarshal(buf.Bytes(), &cnf); err != nil {
		log.Printf("yaml: %s", err)
		return nil
	}

	log.Printf("%#v", cnf)
	return cnf
}
