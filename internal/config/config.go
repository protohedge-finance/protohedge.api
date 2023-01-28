package config

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env    string `yaml:"env"`
	RpcUrl string `yaml:"rpcUrl",envconfig:"RPC_URL"`
	Redis  struct {
		Host     string `yaml:"host",envconfig:"REDIS_HOST"`
		Username string `yaml:"username",envconfig:"REDIS_USERNAME"`
		Password string `yaml:"password",envconfig:"REDIS_PASSWORD"`
	} `yaml:"redis"`
	RebalanceInterval   string `yaml:"rebalanceInterval",envconfig:"REBALANCE_INTERVAL"`
	EventbridgeRuleName string `yaml:"eventbridgeRuleName",envconfig:"EVENTBRIDGE_RULE_NAME"`
	AwsRegion           string `yaml:"awsRegion",envconfig:"AWS_REGION"`
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
	default:
		return Local
	}
}

func NewConfig() *Config {
	env := GetEnvFromEnvVar()
	configFile := GetConfigFileFromEnv(env)
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	f, err := os.Open(filepath.Join(dir, "internal", "config", configFile))

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
