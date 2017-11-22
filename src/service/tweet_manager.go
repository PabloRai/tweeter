package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets []domain.Tweet
var id int

func PublishTweet(twit *domain.Tweet) (int, error) {
	if twit.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if twit.Text == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(twit.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	id++
	twit.Id = id
	tweets = append(tweets, *twit)

	return id, nil
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

func GetTweetById(idTweet int) *domain.Tweet {
	if idTweet <= id {
		fmt.Println(idTweet)
		return &tweets[idTweet-1]
	}
	return nil
}
