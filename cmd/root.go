package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/refs/j/pkg/config"
	"github.com/refs/j/pkg/journal"
	"github.com/spf13/cobra"
)

var jrnl *journal.J

func init() {
	// prefill the file with template.
	tmpl, err := template.New("header").Parse("date:\t{{.Date}}\n------\n\n\n")
	if err != nil {
		log.Fatal(err)
	}

	// TODO use functional options to initialize journal.
	jrnl = &journal.J{
		Config: &config.Config{
			Editor: "vim",
			// $HOME/.j_entries
			Home: fmt.Sprintf("%v/%v", os.Getenv("HOME"), ".j_entries"),
			Format: &config.Format{
				// YYYY-MM-DD
				Date:     time.Now().Format("2006-01-02"),
				Template: tmpl,
			},
		},
	}

	// TODO move this to J's methodset?
	jrnl.Config.FileName = fmt.Sprintf("%v/%v", jrnl.Config.Home, jrnl.Config.Format.Date)

	// ensure there is a J home dir.
	initHome(jrnl)
}

var cfg = config.New()

var rootCmd = &cobra.Command{
	Use:   "j",
	Short: "j is a zero config journaling tool.",
	Long:  `j should help you be more organized and hopefully remember more things over time.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := jrnl.Open(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute runs the command
func Execute() {
	rootCmd.Flags().StringVar(&cfg.Home, "home", cfg.Home, "home directory")
	rootCmd.Flags().StringVar(&cfg.Editor, "editor", cfg.Editor, "default editor (must be accesible on your $PATH)")
	rootCmd.Flags().StringVar(&cfg.Log.Level, "level", cfg.Log.Level, "log level")
	rootCmd.Flags().BoolVar(&cfg.Log.Color, "color", cfg.Log.Color, "colored logs")

	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// initHome panics if initialization fails
func initHome(j *journal.J) {
	_, err := os.Open(j.Config.Home)
	if err != nil {
		// since an error can only be of type *PathError, we're sure
		// no directory exists and we therefore need to create one
		// create j's home, ignoring any errors
		fmt.Printf("HOME not found, creating one at %v\n", j.Config.Home)
		err = os.Mkdir(j.Config.Home, os.FileMode(0644))
		if err != nil {
			log.Fatal(err)
		}
	}
}
