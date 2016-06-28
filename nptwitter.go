package main

import "github.com/ChimeraCoder/anaconda"
import "fmt"
import "net/url"
import "time"

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {

	anaconda.SetConsumerKey("PUT CONSUMER KEY HERE")
	anaconda.SetConsumerSecret("PUT CONSUMER SECRET HERE")
	accesstoken := "PUT ACCESS TOKEN HERE"
	accesssecret := "PUT ACCESS SECRET HERE"
	api := anaconda.NewTwitterApi(accesstoken, accesssecret)
	// u, err := url.Parse("https://api.twitter.com/1.1/statuses/update.json")

	// errCheck(err)
	// TODO: Read in a file of tweets and tweet them every couple of minutes

	//  TODO: Every time a tweet has been made update a counter and make sure that counter is never greater than 180

	v := url.Values{}

	// keywords here
	v.Set("track", "bitcoin, cryptocurrency, blockchaining, blockchain, semantic security, selinux, kali linux, penetration testing, pentesting")

	// start the stream
	stream := api.PublicStreamFilter(v)
	for tweet := range stream.C {
		fmt.Println("------------Task Started------------")
		fmt.Println("Waiting...")
		// sleep for X seconds
		time.Sleep(15 * time.Second)

		fmt.Println("Retweeting...")
		// retweet the tweet
		api.Retweet(tweet.(anaconda.Tweet).Id, false)

		fmt.Println("Waiting again...")
		// wait for X seconds then follow that user
		time.Sleep(15 * time.Second)

		fmt.Println("Following...")
		_, err := api.FollowUser(tweet.(anaconda.Tweet).User.ScreenName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("------------Task Complete------------")
	}
}
