package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

type ImageTweet struct {
	User  string
	Text  string
	Image string
	Date  *time.Time
	Id    int
}

type QuoteTweet struct {
	User  string
	Text  string
	Quote *TextTweet
	Date  *time.Time
	Id    int
}
type Tweet interface {
	PrintableTweet() string
	GetID() int
	GetUser() string
	GetText() string
	SetID(int)
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)

}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.Image)

}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"@%s: \"", tweet.User, tweet.Text, tweet.Quote.User, tweet.Quote.Text)

}

func (tweet *QuoteTweet) GetID() int {
	return tweet.Id
}

func (tweet *TextTweet) GetID() int {
	return tweet.Id
}

func (tweet *ImageTweet) GetID() int {
	return tweet.Id
}

func (tweet *QuoteTweet) SetID(ide int) {
	tweet.Id = ide
}

func (tweet *TextTweet) SetID(ide int) {
	tweet.Id = ide
}

func (tweet *ImageTweet) SetID(ide int) {
	tweet.Id = ide
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func NewTextTweet(user, text string) *TextTweet {
	date := time.Now()
	var id int
	tweet := TextTweet{
		user,
		text,
		&date,
		id,
	}
	return &tweet
}

func NewImageTweet(user, text, url_img string) *ImageTweet {
	date := time.Now()
	var id int
	tweet := ImageTweet{
		user,
		text,
		url_img,
		&date,
		id,
	}
	return &tweet
}

func NewQuoteTweet(user, text string, quote *TextTweet) *QuoteTweet {
	date := time.Now()
	var id int
	tweet := QuoteTweet{
		user,
		text,
		quote,
		&date,
		id,
	}
	return &tweet
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}
