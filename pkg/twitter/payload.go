package twitter

type TweetData struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type Meta struct {
	OldestId    string `json:"oldest_id"`
	NewestId    string `json:"newest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type APIResponse struct {
	Data []TweetData `json:"data"`
	Meta Meta        `json:"meta"`
}
