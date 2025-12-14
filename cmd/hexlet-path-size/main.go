package main

import (
	"context"
	"fmt"

	"os"

	code "github.com/pavloveone/go-project-242"
	"github.com/urfave/cli/v3"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "recursive",
		Aliases:     []string{"r"},
		Usage:       "recursive size of directories",
		DefaultText: "false",
	},
	&cli.BoolFlag{
		Name:        "human",
		Aliases:     []string{"H"},
		Usage:       "human-readable size",
		DefaultText: "false",
	}, &cli.BoolFlag{
		Name:        "all",
		Aliases:     []string{"a"},
		Usage:       "include hidden files and directories",
		DefaultText: "false",
	},
}

func main() {
	command := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: flags,
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path/file is required")
			}

			path := c.Args().First()
			human, all, recursive := c.Bool("human"), c.Bool("all"), c.Bool("recursive")
			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}
			out := fmt.Sprintf("%s\t%s", size, path)
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
