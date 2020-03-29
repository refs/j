package journal

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/refs/j/pkg/config"
)

// Header are optional template header fields.
type Header struct {
	Date  string // entry date.
	Count int    // entry number.
}

// J implements the journaling interface.
type J struct {
	Config *config.Config
}

// Open implements the Journaling interface.
// opens J.Config.File on $EDITOR and waits for the proccess to end.
func (j *J) Open() error {
	cmd := j.cmd()

	j.Config.Format.Template.Execute(
		ensureEntry(j.Config.FileName),
		Header{
			Date: j.Config.Format.Date,
		},
	)

	return startAndWait(cmd)
}

func (j *J) cmd() *exec.Cmd {
	// calls $Editor with the current filename as first argument.
	// $ vim 2020-03-23.txt
	cmd := exec.Command(j.Config.Editor, j.Config.FileName)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd
}

// ensureEntry ensures there is an entry available for today.
func ensureEntry(filename string) *os.File {
	// if there is an entry already, open the editor.
	f, err := os.Open(filename)
	if err != nil {
		// TODO this SHOULD be a debug message.
		fmt.Println(`creating new entry for today`)
		f, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	return f
}

func startAndWait(c *exec.Cmd) error {
	if err := c.Start(); err == nil {
		if err := c.Wait(); err != nil {
			return err
		}
	} else {
		return err
	}

	log.Printf("changes saved.")
	return nil
}
