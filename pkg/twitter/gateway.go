package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Gateway struct {
	config TwitterConfig
	client *http.Client
}

func NewGateway(config TwitterConfig, client *http.Client) *Gateway {
	return &Gateway{config: config, client: client}
}

func (g Gateway) GetUserTimeline() (APIResponse, error) {
	userTimelinePath := "2/users/1048233506470612993/tweets"
	request, err := http.NewRequest(http.MethodGet, g.config.URL+userTimelinePath, nil)
	if err != nil {
		return APIResponse{}, fmt.Errorf("failed to create request, err: %v", err)
	}
	request.Header.Add("Authorization", "Bearer "+g.config.BearerToken)

	res, err := g.client.Do(request)
	if err != nil {
		return APIResponse{}, fmt.Errorf("failed to make request to twitter api, err %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("uh oh, err: %v", err)
	}

	var apiRes APIResponse
	err = json.NewDecoder(res.Body).Decode(&apiRes)
	if err != nil {
		return APIResponse{}, err
	}

	return apiRes, nil
}
