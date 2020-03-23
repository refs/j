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
func (j *J) Open() {
	cmd := j.cmd()

	f := ensureEntry(j.Config.FileName)

	j.Config.Format.Template.Execute(f, Header{
		Date: j.Config.Format.Date,
	})

	if err := cmd.Start(); err == nil {
		if err := cmd.Wait(); err != nil {
			log.Printf("error while editing. Error: %v\n", err)
		}
	} else {
		log.Fatal(err)
	}

	log.Printf("changes saved.")
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
	// if there is an entry already, open the editor in append mode.
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
