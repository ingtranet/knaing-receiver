package receiver

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
)

var logger zerolog.Logger

func newConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetDefault("log_level", "info")

	envs := []string{
		"nats_url",
		"channel_secret",
		"stan_cluster_id",
		"stan_channel",
		"stan_log_channel",
		"log_level",
		"port",
	}

	for _, e := range envs {
		if err := config.BindEnv(e); err != nil {
			return nil, errors.Wrap(err, "binding env failed: " + e)
		}
	}
	for _, env := range envs {
		if !config.IsSet(env) {
			return nil, errors.New("env is not set: " + env)
		}
	}
	return config, nil
}

func newLogger(config *viper.Viper) zerolog.Logger {
	logLevelStr := config.GetString("log_level")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != err {
		panic(errors.Wrap(err, "parsing log_level failed"))
	}

	return zerolog.New(os.Stdout).With().Timestamp().Caller().Str("source", "knaing-receiver").Logger().Level(logLevel)
}

