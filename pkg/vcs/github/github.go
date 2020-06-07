package github

import (
	"fmt"
	"os"
)

func init() {
	// load J_OAUTH_TOKEN from env
	token := os.Getenv("J_OAUTH_TOKEN")
	fmt.Printf("token read from environment:\t%v\n", token)
}

// VCS provides a github implementation of vcs.Versioner.
type VCS struct{}

// Create implements the Versioner interface.
func (v *VCS) Create(entry interface{}) error {
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
