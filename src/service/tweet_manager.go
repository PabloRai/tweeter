package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var id int

type TweetManager struct {
	tweets      map[string][]*domain.Tweet
	userFollows map[string][]string
	twits       []*domain.Tweet
}

func (tweetManager *TweetManager) PublishTweet(twit *domain.Tweet) (int, error) {
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
		tweetManager.tweets[twit.User] = make([]*domain.Tweet, 0)
	}
	tweetManager.tweets[twit.User] = append(userList, twit)

	return id, nil
}

func (tweetManger *TweetManager) GetTweet() *domain.Tweet {
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
	tweetManager.tweets = make(map[string][]*domain.Tweet)
	tweetManager.userFollows = make(map[string][]string)
	tweetManager.twits = make([]*domain.Tweet, 0)
	id = 0

}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	return tweetManager.twits
}

func (tweetManager *TweetManager) GetTweetById(idTweet int) *domain.Tweet {
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

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
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

func (tweetManager *TweetManager) GetTimeline(myUser string) []*domain.Tweet {
	var timeline []*domain.Tweet
	timeline = make([]*domain.Tweet, 0)
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
