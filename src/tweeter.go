package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	tweeterManager := service.NewTweetManager()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			tweet := c.ReadLine()
			twit := domain.NewTweet(user, tweet)
			_, err := tweeterManager.PublishTweet(twit)
			if err != nil {
				c.Println("There was an error (user can't be empty)")
			}
			c.Print("Tweet sent \n")
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a Tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := tweeterManager.GetTweet()
			if tweet != nil {
				c.Println(tweet.User)
				c.Println(tweet.Text)
				c.Println(tweet.Date)
			} else {
				c.Println("There is no tweets")
			}
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "clearTweet",
		Help: "Clear a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			tweeterManager.ClearTweets()
			c.Print("Tweet deleted \n")
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetCountsByUser",
		Help: "Shows the amount of tweets of one user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")

			user := c.ReadLine()

			tweet := tweeterManager.CountTweetsByUser(user)
			c.Println(tweet)
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Shows the amount of tweets of one user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")

			user := c.ReadLine()

			tweet := tweeterManager.GetTweetsByUser(user)
			c.Println(tweet)
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "followUser",
		Help: "Follow an user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")

			user := c.ReadLine()
			c.Print("Write the username to follow: ")
			userToFollow := c.ReadLine()
			err := tweeterManager.Follow(user, userToFollow)
			if err != nil {
				c.Println(err.Error())
			} else {
				c.Println("Done!")
			}
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "getTimeline",
		Help: "Shows your timeline",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")

			user := c.ReadLine()

			tweets := tweeterManager.GetTimeline(user)
			for _, tweet := range tweets {
				c.Println(tweet.User)
				c.Println(tweet.Text)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "sendMessage",
		Help: "Send a direct message to an user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your username: ")

			user := c.ReadLine()
			c.Print("Write the username to send message: ")
			userTo := c.ReadLine()

			c.Print("Write the message to send: ")
			msg := c.ReadLine()
			tweeterManager.SendDirectMessage(user, userTo, msg)
			c.Print("Message ", msg, " sent to ", userTo)
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "getTopics",
		Help: "Shows the current topics",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			tweets := tweeterManager.GetTopics()
			for _, tweet := range tweets {
				c.Println(tweet)
			}
			return
		},
	})
	shell.Run()
}
