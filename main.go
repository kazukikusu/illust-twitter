package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
	"github.com/days365/illust-twitter/logger"
)

func main() {
	api, err := connectTwitterApi()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// ToDo: 一旦仮置き
	searchResult, _ := api.GetSearch("#test", nil)


	tweets := make([]Tweet, 0)

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

func connectTwitterApi() (*anaconda.TwitterApi, error) {
	// ToDo: ファイルを読み込む形で仮実装
	ac, err := ioutil.ReadFile("./twitterAccount.json")
	if err != nil {
		return nil, err
	}

	var t TwitterAccount
	if err := json.Unmarshal(ac, &t); err != nil {
		return nil, err
	}

	return api, nil
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
