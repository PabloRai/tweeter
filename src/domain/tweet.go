package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user, text string) *Tweet {
	date := time.Now()
	var id int
	tweet := Tweet{
		user,
		text,
		&date,
		id,
	}
	return &tweet
}

func (tweet *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet *Tweet) String() string {
	return tweet.PrintableTweet()
}
