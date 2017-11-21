package service

import (
	"github.com/tweeter/src/domain"
)

var tweet domain.Tweet

func PublishTweet(twit *domain.Tweet) {
	tweet = *twit
}

func GetTweet() domain.Tweet {
	return tweet
}

func ClearTweet() {
	tweet.Date = nil
	tweet.Text = ""
	tweet.User = ""
}
