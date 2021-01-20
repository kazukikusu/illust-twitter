package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/days365/illust-twitter/logger"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	c, err := connectTwitterClient()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// ToDo: 一旦仮置き
	search, _, err := c.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#test",
	})

	if err != nil {
		logger.Error(err.Error())
		return
	}

	tweets := make([]Tweet, 0)

	for _, data := range search.Statuses {
		var mediaUrl string
		if len(data.Entities.Media) != 0 {
			medeaList := data.Entities.Media[0]
			mediaUrl = medeaList.MediaURLHttps
		}

		tweet := Tweet{
			ID:       data.ID,
			UserID:   data.User.ID,
			Text:     data.Text,
			MediaUrl: mediaUrl,
		}

		tweets = append(tweets, tweet)
	}

	// ToDo: 一旦デバッグ用に置いておく
	for _, v := range tweets {
		fmt.Printf("%#v\n", v)
	}
}

func connectTwitterClient() (*twitter.Client, error) {
	// ToDo: ファイルを読み込む形で仮実装
	ac, err := ioutil.ReadFile("./twitterAccount.json")
	if err != nil {
		return nil, err
	}

	var t TwitterAccount
	if err := json.Unmarshal(ac, &t); err != nil {
		return nil, err
	}

	config := oauth1.NewConfig(t.ConsumerKey, t.ConsumerSecret)
	token := oauth1.NewToken(t.AccessToken, t.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	return client, nil
}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

type Tweet struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Text     string `json:"text"`
	MediaUrl string `json:"media_url"`
}

type Tweets []Tweet
