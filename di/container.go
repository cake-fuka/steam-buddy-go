package di

import (
	"github.com/cake-fuka/steam-buddy-go/domain"
	"github.com/cake-fuka/steam-buddy-go/infrastructure/env"
)

type Container map[string]interface{}

func NewContainer() Container {
	return Container{}
}

func (c Container) Service() domain.Service {
	key := "Service"
	if _, ok := c[key]; !ok {
		c[key] = domain.NewService(c.Config().SteamToken, c.Config().SteamID, c.Config().SlackToken, c.Config().SlackID)
	}
	return c[key].(domain.Service)
}

func (c Container) Config() *env.Config {
	key := "Config"
	if _, ok := c[key]; !ok {
		conf, err := env.ReadConfig()
		if err != nil {
			panic(err)
		}
		c[key] = conf
	}
	return c[key].(*env.Config)
}
