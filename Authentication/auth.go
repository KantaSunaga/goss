package Authentication

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"context"
	)


func Create() (tc *http.Client) {
	if os.Getenv("GITHUBTOKEN") == "" {
		log.Fatal("環境変数にgithubのAPItokenを設定してください")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUBTOKEN")},
	)
	tc = oauth2.NewClient(ctx, ts)
	return
}