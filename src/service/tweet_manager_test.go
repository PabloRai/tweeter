package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

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

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}
