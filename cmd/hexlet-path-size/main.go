package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	command := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
	}
	err := command.Run(context.Background(), os.Args)
	if err != nil {
		os.Exit(1)
	}
}
