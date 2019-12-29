package j

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "j",
	Short: "j is a zero config journaling tool.",
	Long:  `j should help you be more organized and hopefully remember more things over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*do things*/
	},
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
