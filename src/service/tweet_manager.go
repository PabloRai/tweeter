package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var id int
var tweets map[string][]*domain.Tweet
var twits []*domain.Tweet

func PublishTweet(twit *domain.Tweet) (int, error) {
	if tweets == nil {
		InitializeService()
	}
	if twit.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if twit.Text == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(twit.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	id++
	twit.Id = id
	userList, response := tweets[twit.User]
	twits = append(twits, twit)
	if response == false {
		tweets[twit.User] = make([]*domain.Tweet, 0)
	}
	tweets[twit.User] = append(userList, twit)

	return id, nil
}

func GetTweet() *domain.Tweet {
	if twits != nil && len(twits) > 0 {
		return twits[len(twits)-1]
	}
	return nil
}

func ClearTweets() {
	tweets = nil
	id = 0
	twits = nil

}

func InitializeService() {
	tweets = make(map[string][]*domain.Tweet)
	twits = make([]*domain.Tweet, 0)
	id = 0

}

func GetTweets() []*domain.Tweet {
	return twits
}

func GetTweetById(idTweet int) *domain.Tweet {
	if idTweet <= id {
		return twits[idTweet-1]
	}
	return nil
}

func CountTweetsByUser(user string) int {
	var counter int

	for _, tweet := range twits {
		if tweet.User == user {
			counter++
		}
	}
	return counter
}

func GetTweetsByUser(user string) []*domain.Tweet {
	userList, ok := tweets[user]
	if ok == false {
		return nil
	}
	return userList
}
