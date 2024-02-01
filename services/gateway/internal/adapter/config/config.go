package config

import (
	"runtime/debug"

	"github.com/caarlos0/env/v10"
)

type (
	Overlay struct {
		App     AppConfig
		Logger  LoggerConfig
		Http    HttpConfig
		Kafka   KafkaConfig
		MariaDb MariaDbConfig
		Redis   RedisConfig
	}
	AppConfig struct {
		Name    string `env:"APP_NAME" envDefault:"gateway"`
		Env     string `env:"APP_ENV" envDefault:"dev"`
		Version string
	}
	HttpConfig struct {
		Host string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
		Port int    `env:"HTTP_PORT" envDefault:"8080"`
	}
	KafkaConfig struct {
		Brokers []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`

		TopicWorkerA string `env:"KAFKA_TOPIC_WORKER_A" envDefault:"worker-a"`
		TopicWorkerB string `env:"KAFKA_TOPIC_WORKER_B" envDefault:"worker-b"`
		TopicWorkerC string `env:"KAFKA_TOPIC_WORKER_C" envDefault:"worker-c"`
	}
	MariaDbConfig struct {
		Host     string `env:"MARIADB_HOST" envDefault:"localhost"`
		Port     int    `env:"MARIADB_PORT" envDefault:"3306"`
		Username string `env:"MARIADB_USERNAME" envDefault:"root"`
		Password string `env:"MARIADB_PASSWORD" envDefault:""`
		Database string `env:"MARIADB_DATABASE" envDefault:"gateway"`
	}
	RedisConfig struct {
		Host     string `env:"REDIS_HOST" envDefault:"localhost"`
		Port     int    `env:"REDIS_PORT" envDefault:"6379"`
		Password string `env:"REDIS_PASSWORD" envDefault:""`
		Database int    `env:"REDIS_DATABASE" envDefault:"0"`
	}
	LoggerConfig struct {
		Level string `env:"LOG_LEVEL" envDefault:"debug"`
	}
)

func LoadConfig() (*Overlay, error) {
	var overlay Overlay
	if err := env.Parse(&overlay); err != nil {
		return nil, err
	}

	// get version from build info
	overlay.App.Version = "unknown"
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return &overlay, nil
	}
	for _, v := range info.Settings {
		if v.Key == "vcs.revision" {
			overlay.App.Version = v.Value
			return &overlay, nil
		}
	}

	return &overlay, nil
}
