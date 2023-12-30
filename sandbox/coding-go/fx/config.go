package main

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}
