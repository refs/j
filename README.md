[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9cf9f4575f6940f3a3383e66c2b56d4b)](https://www.codacy.com/manual/refs_2/j?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=refs/j&amp;utm_campaign=Badge_Grade)
# J - A Hacker Journaling Companion written in Go

Just your friendly hacker companion journal written in Go

## Usage

```console
$ go install github.com/refs/j
$ j
```

that's it.

## Gotchas

It currently supports only VIM.

## Roadmap

-   support for version control
-   support encrypted entries
-   configuration
-   support for nano

## User Journey v0

1.  `$ j`
2.  `j` creates a file in J/HOME and opens the user editor (if no editor is provided in config)
3.  start journaling
4.  when saving the file, j commits it to your configured repository (if any)

## Features

-   throw in some encryption to your pages

-   add some configuration

    -   home folder

    -   version control

    -   distributed storage (integration with distributed database?)

    -   editor config option

    -   github integration

    -   extract H1 from file and append it to the filename

        -   credentials, authentication...
        -   git lib: <https://github.com/src-d/go-git>

    -   if there is already a file with today's date, start the editor in append mode

    -   while listing files, dynamically select one for opening

    -   make template contents configurable

    -   query entries

-   bootstrap (rails like)

-   configurable file structure
