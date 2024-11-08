package internal

import (
	"github.com/ankorstore/yokai/fxhealthcheck"
	"github.com/ankorstore/yokai/orm/healthcheck"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fxhealthcheck.AsCheckerProbe(healthcheck.NewOrmProbe),
	)
}
