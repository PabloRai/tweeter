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

func GetTweet() *domain.Tweet {
	if tweets != nil && len(tweets) > 0 {
		return &tweets[len(tweets)-1]
	}
	return nil
}

func ClearTweets() {
	tweets = nil
	id = 0

}

func InitializeService() {
	tweets = make([]domain.Tweet, 0)
}

func GetTweets() []domain.Tweet {
	return tweets
}

func GetTweetById(idTweet int) *domain.Tweet {
	if idTweet <= id {
		return &tweets[idTweet-1]
	}
	return nil
}

func CountTweetsByUser(user string) int {
	var counter int
	for i := 1; i < len(tweets); i++ {
		if GetTweetById(i).User == user {
			counter++
		}
	}
	return counter
}
