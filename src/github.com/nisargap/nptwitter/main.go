package main

import "github.com/ChimeraCoder/anaconda"
import "fmt"
import "net/url"
import "time"
import "encoding/json"
import "os"

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Config struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
	Keywords       string `json:"keywords"`
	SecondsWait    int    `json:"secondsWait"`
}

// Given a configuration filename this function
// decodes the json file into the structure defined
// in Config.
func GetConfig(filename string) Config {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

func main() {

	config := GetConfig("config.json")
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	accesstoken := config.AccessToken
	accesssecret := config.AccessSecret
	api := anaconda.NewTwitterApi(accesstoken, accesssecret)
	// u, err := url.Parse("https://api.twitter.com/1.1/statuses/update.json")

	// errCheck(err)
	// TODO: Read in a file of tweets and tweet them every couple of minutes

	//  TODO: Every time a tweet has been made update a counter and make sure that counter is never greater than 180
	fmt.Println("Started")
	v := url.Values{}

	// keywords here
	v.Set("track", config.Keywords)

	// start the stream
	stream := api.PublicStreamFilter(v)
	var timeWait time.Duration
	timeWait = time.Duration(config.SecondsWait)
	for tweet := range stream.C {
		fmt.Println("------------Task Started------------")
		fmt.Println("Waiting...")
		// sleep for X seconds
		time.Sleep(timeWait * time.Second)

		fmt.Println("Retweeting...")
		// retweet the tweet
		api.Retweet(tweet.(anaconda.Tweet).Id, false)
		time.Sleep(10 * time.Second)
		fmt.Println("Favoriting...")
		_, err := api.Favorite(tweet.(anaconda.Tweet).Id)
		if err != nil {
			fmt.Println(err)
		}
		// wait for X seconds then follow that user
		//time.Sleep(timeWait * time.Second)
		/*
			fmt.Println("Following...")
			_, err := api.FollowUser(tweet.(anaconda.Tweet).User.ScreenName)
			if err != nil {
				fmt.Println(err)
			} */
		fmt.Println("------------Task Complete------------")
	}
}
