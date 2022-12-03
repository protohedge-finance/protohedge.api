package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env    Env
	RpcUrl string `yaml:"rpcUrl",envconfig:"RPC_URL"`
}

func GetConfigFileFromEnv(env Env) string {
	switch env {
	case Production:
		return "config.production.yml"
	}
	return "config.local.yml"
}

func GetEnvFromEnvVar() Env {
	env := os.Getenv("ENV")

	switch env {
	case "production":
		return Production
	}
	return Local
}

func NewConfig() *Config {
	env := GetEnvFromEnvVar()
	configFile := GetConfigFileFromEnv(env)
	f, err := os.Open(fmt.Sprintf("../../internal/config/%s", configFile))

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

	cfg.Env = env

	return &cfg
}
