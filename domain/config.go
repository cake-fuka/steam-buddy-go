package domain

type Config struct {
	SteamID    string
	SteamToken string
	SlackToken string
	SlackID    string
}

func NewConfig(
	steamID string,
	steamToken string,
	slackToken string,
	slackID string,
) *Config {
	return &Config{
		SteamID:    steamID,
		SteamToken: steamToken,
		SlackToken: slackToken,
		SlackID:    slackID,
	}
}
