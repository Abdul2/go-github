package main

import "golang.org/x/oauth2"
import "github.com/google/go-github/github"
import "context"
import "fmt"
import "os"

var i int = 0

func main() {

	token := os.Getenv("githubtoken")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, err := client.Repositories.ListAll(ctx, "", nil)

	for _, repo := range repos {

		fmt.Println(repo.GetName(), i)
		i++
	}

	if err != nil {

		fmt.Print(err.Error())

	}

}
