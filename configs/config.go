package configs

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string
	Database Database `yaml:"database"`
	Kafka    Kafka    `yaml:"messagebroker`
}

type Database struct {
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"databasename"`
}

type Kafka struct {
	Host     string `yaml:"host"`
	Protocol string `yaml:"protocol"`
	Topic    string `yaml: "topic"`
}

func InitConfig() *Config {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	time.Local = location
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}

	viper.SetConfigName(env)
	viper.AddConfigPath("configs")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Config{
		Env: env,
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return config
}
