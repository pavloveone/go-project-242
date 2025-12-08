package main

import (
	"code/internal/code"
	"context"
	"fmt"

	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	command := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path/file is required")
			}

			path := c.Args().First()
			out, err := code.GetSize(path)
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
