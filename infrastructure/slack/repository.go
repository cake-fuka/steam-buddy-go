package slack

import (
	"net/http"
	"strings"
)

type slackRepository struct {
	token string
	id    string
}

func NewSlackRepository(token, id string) *slackRepository {
	return &slackRepository{
		token: token,
		id:    id,
	}
}

func (slack *slackRepository) PostMessage(message string) error {
	baseURL := "https://slack.com/api/chat.postMessage?channel="
	message = strings.Replace(message, " ", "%20", -1)
	URL := baseURL + slack.id + "&text=" + message + "'&pretty=1"

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+slack.token)

	res, _ := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
