package receiver

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func newConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetDefault("log_level", "info")

	envs := []string{
		"nats_url",
		"channel_secret",
		"stan_cluster_id",
		"stan_client_id",
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

func configureGlobalLogger(config *viper.Viper) {
	logLevelStr := config.GetString("log_level")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != err {
		panic(errors.Wrap(err, "parsing log_level failed"))
	}
	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	log.Logger = log.
		With().
		Str("source", "nriy-receiver").
		Caller().
		Timestamp().
		Logger()
}
