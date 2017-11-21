package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

//func TestPublishedTweetIsSaved(t *testing.T) {
//	tweet := "This is my first tweet"
//	service.PublishTweet(tweet)
//
//	if service.GetTweet() != tweet {
//		t.Error("Expected tweet is", tweet)
//	}
//}

//func TestClearTweet(t *testing.T) {
//	tweet := "Tweet to be erased"
//	service.PublishTweet(tweet)
//	service.ClearTweet()
//	if service.GetTweet() != "" {
//		t.Error("Error tweet is not empty", "")
//	}
//}

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("EXPLOTO TODO")
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be null")
	}
}
