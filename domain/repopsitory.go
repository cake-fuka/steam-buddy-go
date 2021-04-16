package domain

type SlackRepository interface {
	PostMessage(message []string) (error)
}

type SteamRepository interface {
	GetState() ([]string, error)
}
