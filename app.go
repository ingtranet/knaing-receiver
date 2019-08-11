package receiver

import "go.uber.org/fx"

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			newConfig,
			newStanClient,
			newServer,
		),
		fx.Invoke(
			configureGlobalLogger,
			configureRouter,
		),
	)

	return app
}
