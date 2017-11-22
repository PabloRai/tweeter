package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweet domain.Tweet

func PublishTweet(twit *domain.Tweet) error {
	if twit.User == "" {
		return fmt.Errorf("user is required")
	} else if twit.Text == "" {
		return fmt.Errorf("text is required")
	} else if len(twit.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}
	tweet = *twit
	return nil
}

func GetTweet() domain.Tweet {
	return tweet
}

func ClearTweet() {
	tweet.Date = nil
	tweet.Text = ""
	tweet.User = ""
}
