package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// ------
// Dependency injection using structs with interface fields.
// ------

type Logger interface {
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type Log struct {
	Logger
}

type ZapFacade struct {
	*zap.SugaredLogger
}

func (z *ZapFacade) Print(args ...interface{})                 { z.Info(args) }
func (z *ZapFacade) Printf(format string, args ...interface{}) { z.Infof(format, args) }

func main() {
	fmt.Println("Logrus (default)")
	// Default Logrus
	{
		log := Log{logrus.New()}
		log.Info("hi")
		log.Warn("hi")
		log.Error("hi")
	}

	fmt.Println("Logrus (json)")
	{
		log := Log{&logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.JSONFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		}}
		log.Info("hi")
		log.Warn("hi")
		log.Error("hi")
	}

	fmt.Println("UberZap")
	{
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		log := Log{&ZapFacade{logger.Sugar()}}
		log.Info("hi")
		log.Warn("hi")
		log.Error("hi")
	}
}
