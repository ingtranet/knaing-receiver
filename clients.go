package receiver

import (
	"fmt"
	"github.com/labstack/gommon/random"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"github.com/rs/zerolog/log"
)

func newStanClient(config *viper.Viper) (stan.Conn, error) {
	clusterID := config.GetString("stan_cluster_id")
	var clientID string
	hostname, err := os.Hostname()
	if err != nil {
		clientID = hostname
	} else {
		clientID = random.String(10)
	}
	natsURL := config.GetString("nats_url")

	log.Info().Msg(fmt.Sprintf("creating connection with %s %s %s", clusterID, clientID, natsURL))
	client, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL), stan.Pings(3, 20))
	if err != nil {
		return nil, errors.Wrap(err, "creating stan client failed: " + natsURL)
	}
	return client, nil
}
