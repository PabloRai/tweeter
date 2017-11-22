package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets []domain.Tweet

func PublishTweet(twit *domain.Tweet) error {
	if twit.User == "" {
		return fmt.Errorf("user is required")
	} else if twit.Text == "" {
		return fmt.Errorf("text is required")
	} else if len(twit.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}
	tweets = append(tweets, *twit)
	return nil
}

func GetTweet() domain.Tweet {

	return tweets[len(tweets)-1]
}

func ClearTweets() {
	tweets = nil
	InitializeService()
}

func InitializeService() {
	tweets = make([]domain.Tweet, 0)
}

func GetTweets() []domain.Tweet {
	return tweets
}
