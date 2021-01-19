package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
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
