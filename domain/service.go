package domain

import "fmt"

type Service interface {
	RecentCheck() error
	WeeklyCheck() error
}

func NewService(steam SteamRepository, slack SlackRepository) *service {
	return &service{
		steam: steam,
		slack: slack,
	}
}

var gameID string

type service struct {
	steam SteamRepository
	slack SlackRepository
}

func (s *service) RecentCheck() error {
	// gameData [0]id [1]game名 [2]user名
	gameData, err := s.steam.GetState()
	if err != nil {
		return err
	}

	// 前はゲームしていなかった
	if gameID == "" {
		// 今もゲームしてない
		if gameData[0] == "" {
			return nil
		}
		//今はゲームしている
		err := s.slack.PostMessage(gameData[2] + " started a game " + gameData[1] + ".")
		if err != nil {
			return err
		}
		gameID = gameData[0]
		return nil
	}
	// 前はゲームしていた
	// 今はゲームしていない
	if gameData[0] == "" {
		s.slack.PostMessage(gameData[2] + " finished a game.")
		if err != nil {
			return err
		}
		gameID = ""
		return nil
	}
	// 今も同じゲームをしている
	if gameData[0] == gameID {
		return nil
	}
	// 今は別のゲームをしている
	s.slack.PostMessage("changed the game that " + gameData[2] + " has started to " + gameData[1] + ".")
	if err != nil {
		return err
	}
	gameID = gameData[0]
	return nil
}

func (s *service) WeeklyCheck() error {
	// gameData [0]id [1]game名 [2]user名
	gameData, err := s.steam.GetRecentlyPlaedGames(5)
	if err != nil {
		return err
	}

	fmt.Println(gameData)

	return nil
}
