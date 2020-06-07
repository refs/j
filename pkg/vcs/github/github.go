package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	token string
)

func init() {
	// load J_OAUTH_TOKEN from env
	token = os.Getenv("J_OAUTH_TOKEN")
	fmt.Printf("token read from environment:\t%v\n", token)
}

// VCS provides a github implementation of vcs.Versioner.
type VCS struct{}

// Create implements the Versioner interface.
func (v *VCS) Create(entry interface{}) error {
	ctx := context.Background()
	// we will assume a branch `master` exists.
	// as test purposes, commit some dummy data to the target repository.

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fileContent := []byte("This is the content of my file\nand the 2nd line of it")

	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("This is my commit message"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("FirstName LastName"), Email: github.String("user@example.com")},
	}

	_, _, err := client.Repositories.CreateFile(ctx, "refs", "journaling", "bang.md", opts)
	if err != nil {
		return err
	}
	return nil
}

// Update implements the Versioner interface.
func (v *VCS) Update(entry interface{}) error {
	return nil
}

// Delete implements the Versioner interface.
func (v *VCS) Delete(entry interface{}) error {
	return nil
}
