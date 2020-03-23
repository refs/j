package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the last 7 entries on your journal",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	var files []string
	// 	filepath.Walk(HOME, func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if !info.IsDir() {
	// 			files = append(files, info.Name())
	// 		}
	// 		return nil
	// 	})

	// 	for _, v := range files {
	// 		fmt.Printf("- %v\n", v)
	// 	}
	// },
}
