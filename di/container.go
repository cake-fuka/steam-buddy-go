package di

import (
	"github.com/cake-fuka/steam-buddy-go/domain"
	"github.com/cake-fuka/steam-buddy-go/infrastructure/env"
	"github.com/cake-fuka/steam-buddy-go/infrastructure/steam"
)

type Container map[string]interface{}

func NewContainer() Container {
	return Container{}
}

func (c Container) Service() domain.Service {
	key := "Service"
	if _, ok := c[key]; !ok {
		c[key] = domain.NewService(c.Config().SlackToken, c.Config().SlackID, c.Steam())
	}
	return c[key].(domain.Service)
}

func (c Container) Steam() domain.SteamRepository {
	key := "Steam"
	if _, ok := c[key]; !ok {
		c[key] = steam.NewSteamRepository(c.Config().SteamToken, c.Config().SteamID)
	}
	return c[key].(domain.SteamRepository)
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
