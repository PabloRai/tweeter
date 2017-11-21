package service

var tweet string

func PublishTweet(twit string) {
	tweet = twit
}

func GetTweet() string {
	return tweet
}
