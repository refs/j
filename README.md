# J - A Hacker Journaling Companion written in Go

Just your friendly hacker companion journal written in Go

## User Journey v0

1. `$ j`
2. `j` creates a file in J/HOME and opens the user editor (if no editor is provided in config)
3. start journaling
4. when saving the file, j commits it to your configured repository (if any)

## Features

- [ ] throw in some encryption to your pages
- [ ] add some configuration
  - [ ] home folder
  - [ ] version control
  - [ ] distributed storage (integration with distributed database?)
  - [ ] different modes of operation? - not sure about this
  - [ ] editor config option
  - [ ] github integration
    - [ ] credentials, authentication...
    - [ ] git lib: https://github.com/src-d/go-git
- [ ] bootstrap (rails like)
- [ ] configurable file structure
  - [ ] support for go templates
  - [ ] yaml maybe? who knows