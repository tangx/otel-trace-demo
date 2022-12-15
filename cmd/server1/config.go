package main

import "os"

type Config struct {
	AppName    string
	NextServer string
}

func (c *Config) SetDefaults() {
	if c.AppName == "" {
		c.AppName = "Unknown Server"
	}
}

var config *Config

func init() {
	c := &Config{
		AppName:    os.Getenv("AppName"),
		NextServer: os.Getenv("NextServer"),
	}

	c.SetDefaults()
	config = c
}
