package twitter_test

import (
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"personal-blog/pkg/twitter"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("hit twitter api and retrieve my first 5 tweets", func(t *testing.T) {
		t.Skip()
		is := is.New(t)

		someTweets := twitter.APIResponse{
			Data: []twitter.TweetData{{
				Id:   "jahsdka",
				Text: "lasjdsj",
			}, {
				Id:   "jahsdka",
				Text: "lasjdsj",
			}},
			Meta: twitter.Meta{
				OldestId:    "hajsdhkash",
				NewestId:    "lasjdksja",
				ResultCount: 2,
				NextToken:   "lsajdlsd",
			},
		}

		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := json.NewEncoder(w).Encode(someTweets)
			is.NoErr(err)
			w.WriteHeader(http.StatusOK)
		}))

		defer testServer.Close()

		config := twitter.TwitterConfig{
			BearerToken: "fakeToken",
			URL:         testServer.URL,
		}

		gateway := twitter.NewGateway(config, &http.Client{})
		feed, err := gateway.GetUserTimeline()

		t.Log(feed)

		is.NoErr(err)
		is.True(len(feed.Data) != 0)
	})
}
