package main

import (
	"code/internal/code"
	"context"
	"fmt"

	"os"

	"github.com/urfave/cli/v3"
)

var flags = []cli.Flag{&cli.BoolFlag{
	Name:    "human",
	Aliases: []string{"H"},
	Usage:   "human-readable size",
}, &cli.BoolFlag{
	Name:    "all",
	Aliases: []string{"a"},
	Usage:   "include hidden files and directories",
}}

func main() {
	command := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: flags,
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path/file is required")
			}

			path := c.Args().First()
			human := c.Bool("human")
			all := c.Bool("all")
			out, err := code.GetSize(path, human, all)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
	err := command.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
