package Commands

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"goss/Authentication"
	"log"
	"os"
)

func Create(username string, repName string) (err error){
	endpoint := fork(username, repName)
	repository:=clone(endpoint, repName)
	repository.remoteAdd(username, repName)
	return
}

 func clone(endpoint string, repName string) (rw *RepoWrapper) {
	 fmt.Println("Clone開始")
	 path, err := os.Getwd()
	 path = path + "/" + repName
	 repository, err := git.PlainClone(path, true, &git.CloneOptions{
		URL:               endpoint,
	 })
	 if err != nil {
	 	log.Println(err)
	 	log.Fatal("cloneに失敗")
	 }
	 fmt.Println("Clone完了" )
	 rw = &RepoWrapper{repository}
	return rw
 }

 func fork(userName string, repName string) (cloneUrl string){
 	 fmt.Println("Fork開始")
	 ctx := context.Background()
	 tc := Authentication.Create()
	 client := github.NewClient(tc)
	 opt := &github.RepositoryCreateForkOptions{
		 Organization: "",
	 }
	 repository,_,_:= client.Repositories.CreateFork(ctx ,userName,repName, opt)
	 fmt.Println("Fork完了")
	 return  *repository.CloneURL
 }

 type RepoWrapper struct {
 	repo *git.Repository
}

func (rw *RepoWrapper )remoteAdd(username string, repName string) {
	 fmt.Println("Upstreamの追加開始")
	 _, err := rw.repo.CreateRemote(&config.RemoteConfig{
		Name: "upstream",
		URLs: []string{"https://github.com/" + username + "/" + repName + ".git"},
	})
	 if err != nil {
	 	log.Println(err)
	 	log.Fatalln("Upstreamの追加に失敗")
	 }
	fmt.Println("Upstreamの追加完了")
	return
}

