package service_test

import (
	"testing"

	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	tweet := "This is my first tweet"
	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}

func TestClearTweet(t *testing.T) {
	tweet := "Tweet to be erased"
	service.PublishTweet(tweet)
	service.ClearTweet()
	if service.GetTweet() != "" {
		t.Error("Error tweet is not empty", "")
	}
}
