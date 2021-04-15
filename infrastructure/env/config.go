package env

import (
	"fmt"
	"os"
	"strings"
)

func ReadConfig() (*Config, error) {
	conf := &Config{}

	var missed []string
	for _, prop := range []struct {
		Field   *string
		EnvName string
	}{
		{&conf.SteamToken, "STEAM_API_KEY"},
		{&conf.SteamID, "STEAM_ID"},
		{&conf.SlackToken, "SLACK_TOKEN"},
		{&conf.SlackID, "SLACK_ID"},
	} {
		env := os.Getenv(prop.EnvName)
		if env == "" {
			missed = append(missed, prop.EnvName)
		} else {
			*prop.Field = env
		}
	}
	if len(missed) > 0 {
		return nil, fmt.Errorf("required environment variables: %s", strings.Join(missed, " "))
	}

	return conf, nil
}
