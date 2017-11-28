package domain

import (
	"os"
)

type MemoryTweetWriter struct {
	Tweets []Tweeter
}

type FileTweetWriter struct {
	File *os.File
}

type ChannelWriter struct {
	WriteConcret Writer
}

func (channelWriter *ChannelWriter) WriteTweet(tweetsToWrite chan Tweeter, quit chan bool) {
	tweet, open := <-tweetsToWrite

	for open {
		channelWriter.WriteConcret.WriteTweet(tweet)
		tweet, open = <-tweetsToWrite
	}
	quit <- true

}

type Writer interface {
	WriteTweet(Tweeter)
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweet Tweeter) {
	memoryTweetWriter.Tweets = append(memoryTweetWriter.Tweets, tweet)

}

func (fileTweetWriter *FileTweetWriter) WriteTweet(tweet Tweeter) {

	fileTweetWriter.File.Write([]byte(tweet.PrintableTweet() + "\n"))

}
