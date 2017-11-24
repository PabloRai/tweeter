package service

import (
	"fmt"
	"strings"

	"github.com/tweeter/src/domain"
)

var id int

type TweetManager struct {
	tweets           map[string][]*domain.TextTweet
	userFollows      map[string][]string
	twits            []*domain.TextTweet
	words            map[string]int
	messagesReceived map[string][]*domain.Message
	favsUser         map[string][]*domain.TextTweet
}

func (tweetManager *TweetManager) PublishTweet(twit *domain.TextTweet) (int, error) {
	if tweetManager.tweets == nil {
		tweetManager.InitializeService()
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
	userList, response := tweetManager.tweets[twit.User]
	tweetManager.twits = append(tweetManager.twits, twit)
	if response == false {
		tweetManager.tweets[twit.User] = make([]*domain.TextTweet, 0)
	}
	tweetManager.tweets[twit.User] = append(userList, twit)
	palabras := strings.Split(twit.Text, " ")
	for _, word := range palabras {
		amount, ok := tweetManager.words[word]
		if ok == false {
			tweetManager.words[word] = 0
		}
		amount++
		tweetManager.words[word] = amount
	}

	return id, nil
}

func (tweetManger *TweetManager) GetTweet() *domain.TextTweet {
	if tweetManger.twits != nil && len(tweetManger.twits) > 0 {
		return tweetManger.twits[len(tweetManger.twits)-1]
	}
	return nil
}

func (tweetManager *TweetManager) ClearTweets() {
	tweetManager.tweets = nil
	id = 0
	tweetManager.twits = nil
	tweetManager.userFollows = nil

}

func (tweetManager *TweetManager) InitializeService() {
	tweetManager.tweets = make(map[string][]*domain.TextTweet)
	tweetManager.userFollows = make(map[string][]string)
	tweetManager.favsUser = make(map[string][]*domain.TextTweet)
	tweetManager.messagesReceived = make(map[string][]*domain.Message)
	tweetManager.words = make(map[string]int)
	tweetManager.twits = make([]*domain.TextTweet, 0)
	id = 0

}

func (tweetManager *TweetManager) GetTweets() []*domain.TextTweet {
	return tweetManager.twits
}

func (tweetManager *TweetManager) GetTweetById(idTweet int) *domain.TextTweet {
	if idTweet <= id {
		return tweetManager.twits[idTweet-1]
	}
	return nil
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	var counter int

	for _, tweet := range tweetManager.twits {
		if tweet.User == user {
			counter++
		}
	}
	return counter
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.TextTweet {
	userList, ok := tweetManager.tweets[user]
	if ok == false {
		return nil
	}
	return userList
}

func (tweetManager *TweetManager) Follow(myUser, userToFollow string) error {
	_, checkUser := tweetManager.tweets[userToFollow]

	if checkUser == false {
		return fmt.Errorf("There is no username with name %s", userToFollow)
	}
	myUserFollowList, ok := tweetManager.userFollows[myUser]
	if ok == false {
		myUserFollowList = make([]string, 0)
	}
	tweetManager.userFollows[myUser] = append(myUserFollowList, userToFollow)
	return nil
}

func (tweetManager *TweetManager) GetTimeline(myUser string) []*domain.TextTweet {
	var timeline []*domain.TextTweet
	timeline = make([]*domain.TextTweet, 0)
	users := tweetManager.userFollows[myUser]
	for index := 0; index < len(users); index++ {

		timeline = append(timeline, tweetManager.GetTweetsByUser(users[index])...)
	}

	timeline = append(timeline, tweetManager.GetTweetsByUser(myUser)...)
	fmt.Println(timeline)
	return timeline
}

func NewTweetManager() *TweetManager {
	tweetManager := TweetManager{}
	return &tweetManager
}

func (tweetManager *TweetManager) GetTopics() [2]string {
	var max int
	var secondMax int

	var maxWords [2]string
	for word, amount := range tweetManager.words {
		if amount > max {
			max = amount
			maxWords[1] = maxWords[0]
			maxWords[0] = word
		} else if amount > secondMax {
			maxWords[1] = word
			secondMax = amount
		}
	}
	return maxWords
}

func (tweetManager *TweetManager) SendDirectMessage(fromUser, toUser, msg string) {

	message := domain.NewMessage(msg)
	userMessagesReceived, response := tweetManager.messagesReceived[toUser]
	if response == false {
		userMessagesReceived = make([]*domain.Message, 0)
	}
	tweetManager.messagesReceived[toUser] = append(userMessagesReceived, message)
}

func (tweetManager *TweetManager) GetMessagesReceivedByUser(user string) []*domain.Message {
	userMessagesReceived, response := tweetManager.messagesReceived[user]
	if response == false {
		return nil
	}
	return userMessagesReceived
}

func (tweetManger *TweetManager) GetUnreadMessages(user string) []*domain.Message {
	userMessagesReceived := tweetManger.GetMessagesReceivedByUser(user)
	var result []*domain.Message
	result = make([]*domain.Message, 0)
	if userMessagesReceived == nil {
		return nil
	}
	for _, message := range userMessagesReceived {
		if message.GetRead() == false {
			result = append(result, message)
		}
	}
	return result
}

func (tweetManager *TweetManager) ReadMessages(user string) []string {
	var messagesUnread []string
	messagesUnread = make([]string, 0)
	messagesStructUnread := tweetManager.GetUnreadMessages(user)
	if messagesStructUnread == nil {
		messagesUnread = append(messagesUnread, "There is no new messages!")
		return messagesUnread
	}
	for _, msg := range messagesStructUnread {
		messagesUnread = append(messagesUnread, msg.GetText())
		msg.Read()
	}
	return messagesUnread
}

func (tweetManager *TweetManager) RetweetById(user string, idTweet int) error {
	myTweets, response := tweetManager.tweets[user]
	if response == false {
		myTweets = make([]*domain.TextTweet, 0)
	}
	if idTweet > len(tweetManager.twits) {
		return fmt.Errorf("The id doesn't belong to any tweets")
	}
	tweetManager.tweets[user] = append(myTweets, tweetManager.twits[idTweet-1])
	return nil
}

func (tweetManager *TweetManager) AddFavorite(user string, idTweet int) error {
	myFavTweets, response := tweetManager.favsUser[user]
	if response == false {
		myFavTweets = make([]*domain.TextTweet, 0)
	}
	if idTweet > len(tweetManager.twits) {
		return fmt.Errorf("The id doesn't belong to any tweets")
	}
	tweetManager.favsUser[user] = append(myFavTweets, tweetManager.twits[idTweet-1])
	return nil
}

func (tweetManager *TweetManager) GetFavsbyUser(user string) []*domain.TextTweet {
	favTweets, response := tweetManager.favsUser[user]
	if response == false {
		return nil
	}
	return favTweets
}
