package j

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var (
	// HOME represents j's home directory
	// TODO get user's home from environment
	HOME = "/Users/aunger/.j_entries"
)

var rootCmd = &cobra.Command{
	Use:   "j",
	Short: "j is a zero config journaling tool.",
	Long:  `j should help you be more organized and hopefully remember more things over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// date format is UNIX
		today := time.Now().Unix()

		// create HOME folder if it doesn't exist
		_, err = os.Open(HOME)
		if err != nil {
			// since an error can only be of type *PathError, we're sure
			// no directory exists and we therefore need to create one
			// create j's home, ignoring any errors
			fmt.Printf("HOME not found, creating one at %v\n", HOME)
			err = os.Mkdir(HOME, os.FileMode(0777)) // TODO please don't use FFA permissions
			if err != nil {
				log.Fatal(err)
			}
		}

		// create a file with today's date
		entryName := fmt.Sprintf("%v/%v", HOME, today)
		entry, err := os.Create(entryName)
		if err != nil {
			log.Fatal(err)
		}

		// open the user editor from environment

		// Open the user's editor routine
		editor := "vim" // get this from the environment, default to... vim?
		editorCmd := exec.Command(editor, entry.Name())

		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		err = editorCmd.Start()
		if err != nil {
			log.Printf("2")
			log.Fatal(err)
		}
		err = editorCmd.Wait()
		if err != nil {
			log.Printf("Error while editing. Error: %v\n", err)
		} else {
			log.Printf("Successfully edited.")
		}

	},
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
