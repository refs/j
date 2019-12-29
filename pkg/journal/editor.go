package journal

import (
	"log"
	"os"
	"os/exec"
)

// OpenEditor opens user editor
func OpenEditor(f *os.File) {
	// Open the user's editor routine
	editor := "vim" // get this from the environment, default to... vim?
	editorCmd := exec.Command(editor, f.Name())

	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	err := editorCmd.Start()
	if err != nil {
		log.Printf("2")
		log.Fatal(err)
	}
	err = editorCmd.Wait()
	if err != nil {
		log.Printf("error while editing. Error: %v\n", err)
	} else {
		log.Printf("changes saved")
	}
}
