package journal

import (
	"log"
	"os"
	"os/exec"
)

// OpenEditor opens user editor
func OpenEditor(f *os.File) {
	editor := os.Getenv("EDITOR")
	editorCmd := exec.Command(editor, f.Name(), "+")

	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	err := editorCmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = editorCmd.Wait()
	if err != nil {
		log.Printf("error while editing. Error: %v\n", err)
	} else {
		log.Printf("changes saved")
	}
}
