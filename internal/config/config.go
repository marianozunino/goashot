package config

import (
	"flag"

	"go.uber.org/fx"
)

type Config interface {
	AdminPassword() string
}

type config struct {
	adminPassword string
}

var Module = fx.Options(
	fx.Provide(registerConfig),
)

func registerConfig() Config {
	var adminPassword = flag.String("passwd", "admin", "Sets the admin password to access the crawler panel")
	flag.Parse()
	return config{*adminPassword}
}

func (c config) AdminPassword() string {
	return c.adminPassword
}
