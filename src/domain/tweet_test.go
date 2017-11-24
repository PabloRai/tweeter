package domain_test

import (
	"testing"

	"github.com/tweeter/src/domain"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")
	text := tweet.PrintableTweet()

	expected_text := "@grupoesfera: This is my tweet"
	if text != expected_text {
		t.Errorf("The expected test is %s but was %s", expected_text, text)
	}
}
