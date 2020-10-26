package journal

import (
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
func (j *J) Open() error {
	cmd := j.prepareCommand()

	if err := j.Config.Format.Template.Execute(
		ensureEntry(j.Config.EntryName),
		Header{
			Date: j.Config.Format.Date,
		},
	); err != nil {
		return err
	}

	return run(cmd)
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
	f, err := os.Open(filename)
	if err != nil {
		print(`creating today's entry`)
		f, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	return f
}

// run execute command <c> and waits for it to return.
func run(c *exec.Cmd) error {
	if err := c.Start(); err == nil {
		if err := c.Wait(); err != nil {
			return err
		}
	} else {
		return err
	}

	print("changes saved.")
	return nil
}
