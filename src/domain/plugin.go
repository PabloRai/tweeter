package domain

type TweetPlugin interface {
	ExecutePlugin(user string) string
}

type FacebookPlugin struct{}
type GooglePlugin struct{}
type CountPlugin struct{}

func (facebookPlugin *FacebookPlugin) ExecutePlugin(user string) string {
	return "Posted on Facebook"
}

func (googlePlugin *GooglePlugin) ExecutePlugin(user string) string {
	return "Posted on Google"
}

func (countPlugin *CountPlugin) ExecutePlugin(ammount string) string {
	return ammount
}
