package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestFollowUser(t *testing.T) {
	tweetManager := service.NewTweetManager()
	tweet := domain.NewTweet("grupoesfera", "sarasa")
	tweet1 := domain.NewTweet("nportas", "mytw")
	tweet2 := domain.NewTweet("meli", "melitw")
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(tweet1)
	tweetManager.PublishTweet(tweet2)
	err := tweetManager.Follow("grupoesfera", "nportas")
	err1 := tweetManager.Follow("grupoesfera", "meli")
	timeline := tweetManager.GetTimeline("grupoesfera")
	if err != nil || err1 != nil {
		t.Error("Expected no errors")
	}
	if len(timeline) != 3 {
		t.Errorf("The timeline should have three tweets")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
   all over the world. To date all community oriented activities have been organized by the community
   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

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

func TestTrendingTopics(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(user, secondText)
	thirdTweet := domain.NewTweet(anotherUser, text)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	topics := tweetManager.GetTopics()
	if topics[0] != "a" || topics[1] != "b" {
		t.Error("Wrong topics, should be 'a' and 'b'")
	}

}

func TestUserReceiveADirectMessage(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.SendDirectMessage(user, anotherUser, "HOLA")
	anotherUserMessages := tweetManager.GetMessagesReceivedByUser(anotherUser)
	if anotherUserMessages == nil || len(anotherUserMessages) != 1 {
		t.Error("anotherUserMessages should have one unread message")
	}

}

func TestUserGetUnreadMessage(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.SendDirectMessage(user, anotherUser, "HOLA")
	anotherUserMessages := tweetManager.GetUnreadMessages(anotherUser)
	if anotherUserMessages == nil || len(anotherUserMessages) != 1 {
		t.Error("anotherUserMessages should have one unread message")
	}

}

func TestReadMessage(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.SendDirectMessage(user, anotherUser, "HOLA")
	messages := tweetManager.ReadMessages(anotherUser)
	anotherUserMessages := tweetManager.GetUnreadMessages(anotherUser)
	if len(anotherUserMessages) != 0 || messages[0] != "HOLA" {
		t.Error("anotherUserMessages should have no new messages")
	}

}
