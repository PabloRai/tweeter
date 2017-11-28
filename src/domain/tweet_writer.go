package domain

import (
	"os"
)

type MemoryTweetWriter struct {
	Tweets []Tweet
}

type FileTweetWriter struct {
	File *os.File
}

type ChannelWriter struct {
	WriteConcret Writer
}

func (channelWriter *ChannelWriter) WriteTweet(tweetsToWrite chan Tweet, quit chan bool) {
	tweet, open := <-tweetsToWrite

	for open {
		channelWriter.WriteConcret.WriteTweet(tweet)
		tweet, open = <-tweetsToWrite
	}
	quit <- true

}

type Writer interface {
	WriteTweet(Tweet)
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweet Tweet) {
	memoryTweetWriter.Tweets = append(memoryTweetWriter.Tweets, tweet)

}

func (fileTweetWriter *FileTweetWriter) WriteTweet(tweet Tweet) {

	fileTweetWriter.File.Write([]byte(tweet.PrintableTweet() + "\n"))

}
