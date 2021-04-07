package main

import (
	"fmt"
	"os"

	"github.com/zengqiang96/gdraw/cmd/gdraw/command"
)

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "gdraw: %s\n", err)
		os.Exit(1)
	}
}
