package config

import "os"

type Config struct {
    EmailServer string
}

func LoadConfig() *Config {
    return &Config{
        EmailServer: os.Getenv("EMAIL_SERVER"),
    }
}
