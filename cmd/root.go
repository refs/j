package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/refs/j/pkg/journal"
	"github.com/spf13/cobra"
)

var (
	homeName = `.j_entries`
	// HOME represents j's home directory
	HOME = fmt.Sprintf("%v/%v", os.Getenv("HOME"), homeName)
)

type header struct {
	Date  string // entry date
	Count int    // entry number
}

var rootCmd = &cobra.Command{
	Use:   "j",
	Short: "j is a zero config journaling tool.",
	Long:  `j should help you be more organized and hopefully remember more things over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var f *os.File

		// format: YYYY-MM-DD
		today := time.Now().Format("2006-01-02")

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

		// if there is an entry already, open the editor in append mode
		f, err = os.Open(entryName)
		if err != nil {
			fmt.Println(`creating new entry for today`)
			f, err = os.Create(entryName)
			if err != nil {
				log.Fatal(err)
			}

			// prefill the file with the contents of a template (template is ofc configurable)
			tmpl, err := template.New("header").Parse("date:\t{{.Date}}\n------\n\n\n")
			if err != nil {
				log.Fatal(err)
			}

			tmpl.Execute(f, header{
				Date: today,
			})
		}

		journal.OpenEditor(f)
	},
}

// Execute runs the command
func Execute() {
	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
