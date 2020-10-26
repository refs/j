package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"crypto/sha1"
	"path/filepath"

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
func (v *VCS) Commit(path interface{}) error {
	ctx := context.Background()
	// we will assume a branch `master` exists.
	// as test purposes, commit some dummy data to the target repository.

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fileContent, err := ioutil.ReadFile(path.(string))
	if err != nil {
		return err
	}

	base := filepath.Base(path.(string))

	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("revision 2"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("j-bot"), Email: github.String("hello+automation@zyxan.io")},
		SHA: getHash([]byte(base)),
	}

	blob, _, err := client.Git.GetBlob(ctx, "refs", "journaling", base)
	if err != nil && blob.SHA != nil {
		fmt.Printf("\nblob: %v\n", blob.SHA)
		return err
	}

	updateOpts := &github.RepositoryContentFileOptions{
		Message:   github.String("revision x"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("j-bot"), Email: github.String("hello+automation@zyxan.io")},
		SHA: &base,
	}

	_, _, err = client.Repositories.CreateFile(ctx, "refs", "journaling", base, opts)
	if err != nil {
		// attempt to update the file
		_, _, err = client.Repositories.UpdateFile(ctx, "refs", "journaling", base, updateOpts)
		if err != nil {
			return err
		}
	}

	return nil
}

func getHash(e []byte) *string {
	t := ""
	var r = &t
	h := sha1.New()
	h.Write(e)
	*r = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return r
}