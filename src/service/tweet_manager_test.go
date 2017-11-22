package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	service.InitializeService()

	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)
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
	service.InitializeService()

	var tweet *domain.Tweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	var err error
	_, err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublihed(t *testing.T) {
	service.InitializeService()

	var tweet *domain.Tweet
	user := "grupoesfera"
	var text string
	tweet = domain.NewTweet(user, text)
	var err error
	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
		return
	}
}

func TestTweetWhichExceding140CharacterIsNotPublished(t *testing.T) {
	service.InitializeService()

	var tweet *domain.Tweet
	user := "NICO"
	text := `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`

	tweet = domain.NewTweet(user, text)

	var err error

	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	service.InitializeService()

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, &firstPublishedTweet, firstPublishedTweet.Id, user, text) {
		return
	}

	if !isValidTweet(t, &secondPublishedTweet, secondPublishedTweet.Id, user, secondText) {
		return
	}
	service.ClearTweets()

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
	service.ClearTweets()

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}
