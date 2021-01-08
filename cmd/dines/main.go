package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/0n1shi/dines/pkg/disasm"
	cli "github.com/urfave/cli/v2"
)

const (
	Version = "v0.0.1"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "rom",
			Usage:    "A file path of NES ROM",
			Required: true,
		},
	}

	app := cli.App{
		Name:    "Dines",
		Usage:   `A disassembler for customed 8-bit microprocessor, "MOS Technology 6502" in Nintendo Entertainment System written in Golang.`,
		Version: Version,
		Action:  run,
		Flags:   flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	romFile := c.String("rom")

	data, err := ioutil.ReadFile(romFile)
	if err != nil {
		return err
	}

	dines, err := disasm.NewDines()
	if err != nil {
		return err
	}

	result, err := dines.Disassemble(data)
	if err != nil {
		return err
	}

	dines.Dump(result)

	return nil
}
