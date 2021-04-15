package domain

import "fmt"

type Service interface {
	ObservSteam() error
}

func NewService(steamToken, steamID, slackToken, slackID string) Service {
	return &service{
		steamToken: steamToken,
		steamID:    steamID,
		slackToken: slackToken,
		slackID:    slackID,
	}
}

type service struct {
	steamToken string
	steamID    string
	slackToken string
	slackID    string
}

func (s *service) ObservSteam() error {
	fmt.Printf("OK!")
	return nil
}
