package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	RpcUrl string `yaml:"rpcUrl",envconfig:"RPC_URL"`
}

func GetConfigFileFromEnvVar() string {
	envFile := os.Getenv("ENV_FILE")

	switch envFile {
	case "production":
		return "config.production.yml"
	}
	return "config.local.yml"
}

func NewConfig() *Config {
	f, err := os.Open(fmt.Sprintf("../../internal/config/%s", GetConfigFileFromEnvVar()))

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("", &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}
