package domain

type MemoryTweetWriter struct {
	Tweets []Tweeter
}

type FileTweetWriter struct {
	path string
}

type ChannelWriter struct {
	WriteConcret Writer
}

func (channelWriter *ChannelWriter) WriteTweet(chan1 chan Tweeter, status chan bool) {
	channelWriter.WriteConcret.WriteTweet(chan1, status)
}

type Writer interface {
	WriteTweet(tweetsToWrite chan Tweeter, quit chan bool)
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweetsToWrite chan Tweeter, quit chan bool) {
	memoryTweetWriter.Tweets = make([]Tweeter, 0)
	tweet, open := <-tweetsToWrite

	for open {
		memoryTweetWriter.Tweets = append(memoryTweetWriter.Tweets, tweet)
		tweet, open = <-tweetsToWrite
	}
	quit <- true
}
