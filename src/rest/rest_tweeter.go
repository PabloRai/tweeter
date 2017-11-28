package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

type GinTweet struct {
	User string
	Text string
}

type GinServer struct {
	tweetManager *service.TweetManager
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager}
}

func (server *GinServer) StartGinServer() {

	router := gin.Default()

	router.GET("/listTweets", server.listTweets)
	router.GET("/listTweets/:user", server.listTweets)
	router.POST("publishTweet", server.publishTweet)

	go router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {

	c.JSON(http.StatusOK, server.tweetManager.GetTweets())
}

func (server *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	c.JSON(http.StatusOK, server.tweetManager.GetTweetsByUser(user))
}

func (server *GinServer) publishTweet(c *gin.Context) {

	quit := make(chan bool)

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish, quit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
