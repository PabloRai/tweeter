package domain

type MemoryTweetWriter struct {
	Tweets []Tweeter
}

type FileTweetWriter struct {
	path string
}

type Writer interface {
	WriteTweet(tweetsToWrite chan Tweeter, quit chan bool)
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweetsToWrite chan Tweeter, quit chan bool) {
	memoryTweetWriter.Tweets = make([]Tweeter, 0)
	for tweet := range tweetsToWrite {
		memoryTweetWriter.Tweets = append(memoryTweetWriter.Tweets, tweet)
	}
}
