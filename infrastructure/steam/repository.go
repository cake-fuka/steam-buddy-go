package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type steamRepository struct {
	token string
	id    string
}

func NewSteamRepository(token, id string) *steamRepository {
	return &steamRepository{
		token: token,
		id:    id,
	}
}

func (steam *steamRepository) GetState() ([]string, error) {
	var data GameData

	baseURL := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?"
	url := baseURL + "key=" + steam.token + "&steamids=" + steam.id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []string{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return []string{}, err
	}

	return []string{data.Response.Players[0].Gameid, data.Response.Players[0].Gameextrainfo, data.Response.Players[0].Personaname}, nil
}
