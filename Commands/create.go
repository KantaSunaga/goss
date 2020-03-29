package Commands

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"goss/Authentication"
	"goss/Colors"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"io"
	"log"
	"os"
)

func Create(username string, repName string) (err error){
	endpoint, err := fork(username, repName)
	if err != nil{
		log.Fatal(Colors.Red("Fork faild"))
	}
	clone(endpoint)
	return
}

 func clone(endpoint string){
	 fs := memfs.New()
	 storer := memory.NewStorage()
	 _, err := git.Clone(storer, fs, &git.CloneOptions{
		 URL: endpoint,
	 })
	 if err != nil {
	 	fmt.Println("エラーだよ")
	 	log.Fatal(err)
	 }
	 changelog, err := fs.Open("CHANGELOG")
	 if err != nil {
		 fmt.Println("エラーだよ1")
		 log.Fatal(err)
	 }

	 io.Copy(os.Stdout, changelog)
 }

 func fork(userName string, repName string) (cloneUrl string ,err error){
	 ctx := context.Background()
	 tc := Authentication.Create()
	 client := github.NewClient(tc)
	 opt := &github.RepositoryCreateForkOptions{
		 Organization: "",
	 }
	 repository,_,_:= client.Repositories.CreateFork(ctx ,userName,repName, opt)
	 fmt.Println(*repository.ForksURL)
	 return  *repository.CloneURL, err
 }

func remoteAdd(user string, repo string) {


}

func sample() {

}
