package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/urfave/cli/v2"
)

func TopN(n int, r io.Reader) ([]int64, error) {
	top := make([]int64, 0, n)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		curr, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			continue
		}

		if len(top) == 0 {
			top = append(top, curr)
			continue
		}

		if curr < top[len(top)-1] {
			continue
		}

		for i := 0; i < len(top); i += 1 {
			if curr < top[i] {
				continue
			}

			top = slices.Insert(top, i, curr)
			break
		}

		if len(top) > n {
			top = top[:n]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return top, nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "Path to file to scan",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Usage:   "Number of top lines to return",
				Value:   10,
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Int("number") < 1 {
				return fmt.Errorf("number must be greater than 0")
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			f, err := os.Open(ctx.String("file"))
			if err != nil {
				return err
			}
			defer f.Close()

			top, err := TopN(ctx.Int("number"), f)
			if err != nil {
				return err
			}

			for _, entry := range top {
				fmt.Println(entry)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
