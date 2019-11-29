package receiver

import "go.uber.org/fx"

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			newConfig,
			newStanClient,
			newServer,
			newLogger,
		),
		fx.Populate(
			&logger,
		),
		fx.Invoke(
			configureRouter,
		),
	)

	return app
}
