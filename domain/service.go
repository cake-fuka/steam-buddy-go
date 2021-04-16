package domain

import (
	"fmt"
)

type Service interface {
	ObservSteam() error
}

func NewService(slackToken, slackID string, steam SteamRepository) *service {
	return &service{
		slackToken: slackToken,
		slackID:    slackID,
		steam:      steam,
	}
}

var gameID string

type service struct {
	slackToken string
	slackID    string
	steam      SteamRepository
}

func (s *service) ObservSteam() error {
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
		fmt.Println(gameData[2] + " started a game「" + gameData[1] + "」.")
		gameID = gameData[0]
		return nil
	}
	// 前はゲームしていた
	// 今はゲームしていない
	if gameData[0] == "" {
		fmt.Println(gameData[2] + " finished a game.")
		gameID = ""
		return nil
	}
	// 今も同じゲームをしている
	if gameData[0] == gameID {
		return nil
	}
	// 今は別のゲームをしている
	fmt.Println("changed the game that " + gameData[2] + " has started to「" + gameData[1] + "」.")
	gameID = gameData[0]
	return nil
}
