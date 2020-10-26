package fs

import (
	"fmt"
	"github.com/refs/j/pkg/config"
	vcs "github.com/refs/j/pkg/vcs/github"
	"log"
	"os"
	"os/exec"
)

// Header are optional template header fields.
type Header struct {
	Date  string // entry date.
	Count int    // entry number.
}

// J implements the Journaling interface.
type J struct {
	Config *config.Config
}

// Open implements the Journaling interface.
func (j *J) Open() error {
	cmd := j.prepareCommand()

	err := j.writeFileHeader()
	if err != nil {
		return err
	}

	return j.run(cmd)
}

func (j *J) writeFileHeader() error {
	if !entryExists(j.Config.EntryName) {
		if err := j.Config.Format.Template.Execute(
			ensureEntry(j.Config.EntryName),
			Header{
				Date: j.Config.Format.Date,
			},
		); err != nil {
			return err
		}
	}

	return nil
}

func entryExists(e string) bool {
	if _, err := os.Open(e); err != nil {
		return false
	}
	return true
}


// Open implements the Journaling interface.
func (j *J) Close() error {
	return fmt.Errorf("not implemented")
}

// prepareCommand prepares a command to be executed on the configured editor and entry file.
func (j J) prepareCommand() *exec.Cmd {
	cmd := exec.Command(j.Config.Editor, j.Config.EntryName)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd
}

// ensureEntry ensures there is an entry available for today.
func ensureEntry(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// run execute command <c> and waits for it to return.
func (j *J) run(c *exec.Cmd) error {
	if err := c.Start(); err == nil {
		if err := c.Wait(); err != nil {
			return err
		}

		// start commit routine and commit changes to vsc.
		if err := j.commitChanges(); err != nil {
			return err
		}
	} else {
		return err
	}

	print("changes saved.")
	return nil
}

func (j J) commitChanges() error {
	vc := vcs.VCS{}

	return vc.Commit(j.Config.EntryName)
}