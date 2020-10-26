package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the last 7 entries on your journal",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	var files []string
	if err := filepath.Walk(j.Config.Home, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		fmt.Printf("- %v\n", v)
	}
}