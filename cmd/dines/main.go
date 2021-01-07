package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "rom",
			Usage:    "file path of NES ROM",
			Required: true,
		},
	}

	app := cli.App{
		Name:   "Dines",
		Usage:  `A disassembler for customed 8-bit microprocessor, "MOS Technology 6502" in Nintendo Entertainment System written in Golang.`,
		Action: run,
		Flags:  flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	filePath := c.String("rom")

	fmt.Printf("file path: %s\n", filePath)
	return nil
}
