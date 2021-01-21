package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/days365/illust-twitter/logger"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	tc, err := connectTwitterClient()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// ToDo: 一旦仮置き
	search, _, err := tc.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#test",
	})

	if err != nil {
		logger.Error(err.Error())
		return
	}

	ctx := context.Background()
	dc, err := connectDatastoreClient(ctx)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	tweets, keys := setTweets(search, make([]Tweet, 0))

	if _, err := dc.PutMulti(ctx, keys, tweets); err != nil {
		logger.Error(err.Error())
		return
	}
}

func setTweets(search *twitter.Search, tweets []Tweet) ([]Tweet, []*datastore.Key) {
	keys := make([]*datastore.Key, 0)

	for _, data := range search.Statuses {
		var mediaUrl string
		if len(data.Entities.Media) != 0 {
			medeaList := data.Entities.Media[0]
			mediaUrl = medeaList.MediaURLHttps
		}

		tweet := Tweet{
			ID:         data.ID,
			UserID:     data.User.ID,
			ScreenName: data.User.ScreenName,
			Text:       data.Text,
			MediaUrl:   mediaUrl,
			CreatedAt:  data.CreatedAt,
			InsertedAt: time.Now(),
		}

		tweets = append(tweets, tweet)

		key := datastore.IDKey(getKind(), tweet.ID, nil)
		keys = append(keys, key)
	}

	return tweets, keys
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

func connectDatastoreClient(ctx context.Context) (*datastore.Client, error) {
	dc, err := datastore.NewClient(ctx, getProject())
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return dc, nil
}

func getProject() string {
	return os.Getenv("DATASTORE_PROJECT_ID")
}

func getKind() string {
	return os.Getenv("DATASTORE_KIND")
}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

type Tweet struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	ScreenName string    `json:"screen_name"`
	Text       string    `json:"text"`
	MediaUrl   string    `json:"media_url"`
	CreatedAt  string    `json:"created_at"`
	InsertedAt time.Time `json:"inserted_at"`
}

type Tweets []Tweet
