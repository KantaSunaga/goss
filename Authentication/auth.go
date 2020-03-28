package Authentication

import (
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"context"
	)


func Create() (tc *http.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUBTOKEN")},
	)
	tc = oauth2.NewClient(ctx, ts)
	return
}