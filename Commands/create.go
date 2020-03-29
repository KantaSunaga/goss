package Commands

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"goss/Authentication"
	"goss/Colors"
	"log"
	"os"
)

func Create(username string, repName string) (err error){
	endpoint, err := fork(username, repName)
	if err != nil{
		fmt.Println("ここですお111")

		log.Fatal(Colors.Red("Fork faild"))
	}
	repository := clone(endpoint, repName)
	repository.remoteAdd(username, repName)
	fmt.Println("ここまで来た3")
	return
}

 func clone(endpoint string, repName string) (repository *git.Repository) {
	 path, err := os.Getwd()
	 path = path + "/" + repName
	 fmt.Println(path)
	 fmt.Println("ここまで来た")
	 repository, err = git.PlainClone(path, true, &git.CloneOptions{
		URL:               endpoint,
	 })
	 fmt.Println("ここまで来た33333")
	 if err != nil {
	 	fmt.Println("エラーだよ")
	 	log.Fatal(err)
	 }
	 fmt.Println("ここまで来た2")
	return repository
 }

 func fork(userName string, repName string) (cloneUrl string ,err error){
	 ctx := context.Background()
	 tc := Authentication.Create()
	 client := github.NewClient(tc)
	 opt := &github.RepositoryCreateForkOptions{
		 Organization: "",
	 }
	 repository,_,_:= client.Repositories.CreateFork(ctx ,userName,repName, opt)
	 return  *repository.CloneURL, err
 }

func (r *git.Repository )remoteAdd(username string, repName string) {
	 _, err := r.CreateRemote(&config.RemoteConfig{
		Name: "upstream",
		URLs: []string{"https://github.com/" + username + "/" + repName + ".git"},
	})
	 if err != nil {
	 	fmt.Printf("remote追加中にエラー")
	 }
	return
}

func sample() {

}
