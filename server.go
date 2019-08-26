package receiver

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	bot "github.com/line/line-bot-sdk-go/linebot"
	"go.uber.org/fx"
	"github.com/rs/zerolog/log"
)

func newServer(lc fx.Lifecycle, config *viper.Viper) (*echo.Echo, error) {
	server := echo.New()
	server.Use(middleware.Logger())
	port := config.GetString("port")

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.Start(":" + port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server, nil
}

func configureRouter(server *echo.Echo, config *viper.Viper, client stan.Conn) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	channelSecret := config.GetString("channel_secret")
	stanChannel := config.GetString("stan_channel")
	server.POST("/", func(c echo.Context) error {
		events, err := bot.ParseRequest(channelSecret, c.Request())
		if err != nil {
			return errors.Wrap(err, "body parsing failed")
		}
		for _, event := range events {
			b, err := event.MarshalJSON()
			if err != nil {
				return errors.Wrap(err, "json marshaling failed")
			}
			log.Debug().Msg("dealing with: " + string(b))
			err = client.Publish(stanChannel, b)
			if err != nil {
				return errors.Wrap(err, "stan publish failed")
			}
		}
		return nil
	})
}
