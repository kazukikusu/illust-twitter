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

	tweets := make([]*Tweet, 0)

	// 取得対象：tweetのurlとなるもの。画像URL or データ
	for _, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.ID = data.Id
		tweet.Text = data.Text
		tweet.UserID = data.User.Id

		if len(data.Entities.Media) != 0 {
			medeaList := data.Entities.Media[0]
			tweet.MediaUrl = medeaList.Media_url_https
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
	raw, error := ioutil.ReadFile("./twitterAccount.json")
	if error != nil {
		return nil, error
	}

	var t TwitterAccount
	json.Unmarshal(raw, &t)
	api := anaconda.NewTwitterApiWithCredentials(t.AccessToken, t.AccessTokenSecret, t.ConsumerKey, t.ConsumerSecret)

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

type Tweets *[]Tweet
