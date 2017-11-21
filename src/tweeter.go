package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your tweet: ")
			tweet := c.ReadLine()
			service.PublishTweet(tweet)
			c.Print("Tweet sent \n")
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a Tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := service.GetTweet()
			c.Println(tweet)
			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "clearTweet",
		Help: "Clear a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			service.ClearTweet()
			c.Print("Tweet deleted \n")
			return
		},
	})
	shell.Run()
}
