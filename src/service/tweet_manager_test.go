/* package service_test

import (
	"testing"






















	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

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
	tweet := domain.NewTextTweet("grupoesfera", "sarasa")
	tweet1 := domain.NewTextTweet("nportas", "mytw")
	tweet2 := domain.NewTextTweet("meli", "melitw")
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

	var tweet *domain.TextTweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

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

	var tweet *domain.TextTweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
   all over the world. To date all community oriented activities have been organized by the community
   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

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

	var tweet, secondTweet *domain.TextTweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

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

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

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

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

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

func isValidTweet(t *testing.T, tweet *domain.TextTweet, id int, user, text string) bool {

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

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(user, secondText)
	thirdTweet := domain.NewTextTweet(anotherUser, text)
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

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(anotherUser, secondText)
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

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(anotherUser, secondText)
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

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.SendDirectMessage(user, anotherUser, "HOLA")
	messages := tweetManager.ReadMessages(anotherUser)
	anotherUserMessages := tweetManager.GetUnreadMessages(anotherUser)
	if len(anotherUserMessages) != 0 || messages[0] != "HOLA" {
		t.Error("anotherUserMessages should have no new messages")
	}

}

func TestRetweet(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	err := tweetManager.RetweetById(user, 2)
	myTweets := tweetManager.GetTweetsByUser(user)
	if len(myTweets) != 2 || err != nil {
		t.Error("The tweets in myTweets should be 2")
	}

}

func TestFavorites(t *testing.T) {
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"
	anotherUser := "nick"
	text := "a b c e f"
	secondText := "a a b c d"

	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(anotherUser, secondText)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	err := tweetManager.AddFavorite(user, 2)
	myFavTweets := tweetManager.GetFavsbyUser(user)
	if len(myFavTweets) != 1 || err != nil {
		t.Error("The tweets in myFavTweets should be 1")
	}

}

func TestShareOnGoogle(t *testing.T) {
	tweetManager := service.NewTweetManager()
	err := tweetManager.AddPlugin(&domain.GooglePlugin{})
	plugins := tweetManager.ExecutePlugins("user")
	if plugins[0] != "Posted on Google" || err != nil {
		t.Error("Plugin for google should be active")
	}
}

func TestShareOnFacebook(t *testing.T) {
	tweetManager := service.NewTweetManager()
	err := tweetManager.AddPlugin(&domain.FacebookPlugin{})
	plugins := tweetManager.ExecutePlugins("user")
	if plugins[0] != "Posted on Facebook" || err != nil {
		t.Error("Plugin for google should be active")
	}

}

func TestCountTweetsPlugin(t *testing.T) {
	tweetManager := service.NewTweetManager()
	err := tweetManager.AddPlugin(&domain.CountPlugin{})
	plugins := tweetManager.ExecutePlugins("user")
	if plugins[0] != "user" || err != nil {
		t.Error("Plugin for google should be active")
	}

}

func TestCantAddTwoEqualsPlugins(t *testing.T) {

	tweetManager := service.NewTweetManager()
	err1 := tweetManager.AddPlugin(&domain.CountPlugin{})
	err2 := tweetManager.AddPlugin(&domain.CountPlugin{})
	if err1 != nil || err2 == nil || err2.Error() != "Plugin already exists" {
		t.Errorf("Error should appear ")
	}

}

func TestCanWriteATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "Async tweet")
	tweet2 := domain.NewTextTweet("grupoesfera", "Async tweet2")

	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetsToWrite := make(chan domain.Tweeter)
	quit := make(chan bool)

	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2
	close(tweetsToWrite)

	<-quit

	// Validation
	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

	if memoryTweetWriter.Tweets[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected")
	}
}
*/

package service_test

import (
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet, quit)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, (*publishedTweet), id, user, text)

	<-quit

	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

}

/* func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet, quit)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet, quit)

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
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
    all over the world. To date all community oriented activities have been organized by the community
    with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet, quit)

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
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	quit := make(chan bool)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet, quit)
	secondId, _ := tweetManager.PublishTweet(secondTweet, quit)

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
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet, quit)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	quit := make(chan bool)

	tweetManager.PublishTweet(tweet, quit)
	tweetManager.PublishTweet(secondTweet, quit)
	tweetManager.PublishTweet(thirdTweet, quit)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	quit := make(chan bool)

	firstId, _ := tweetManager.PublishTweet(tweet, quit)
	secondId, _ := tweetManager.PublishTweet(secondTweet, quit)
	tweetManager.PublishTweet(thirdTweet, quit)

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
*/

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) bool {

	if tweet.GetID() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetID())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.GetUser(), tweet.GetText())
		return false
	}

	/* if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	} */

	return true

}
