package receiver

import (
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func newStanClient(config *viper.Viper) (stan.Conn, error) {
	clusterID := config.GetString("stan_cluster_id")
	clientID := config.GetString("stan_client_id")
	natsURL := config.GetString("nats_url")

	client, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return nil, errors.Wrap(err, "creating stan client failed: " + natsURL)
	}
	return client, nil
}
