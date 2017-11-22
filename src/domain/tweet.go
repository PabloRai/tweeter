package domain

import (
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
