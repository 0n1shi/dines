package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/0n1shi/dines/pkg/disasm"
	cli "github.com/urfave/cli/v2"
)

const (
	version = "unknown"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "rom",
			Usage:    "A file path of NES ROM",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "output",
			Usage:       "output format, \"json\" or \"yaml\", default is like a typical diassembler",
			DefaultText: "normal",
		},
		&cli.BoolFlag{
			Name:  "color",
			Usage: "color output (*only available without \"output\" option)",
		},
		&cli.IntFlag{
			Name:  "max",
			Usage: "max  number of lines of output excluding header information (*only available without \"output\" option)",
			Value: -1,
		},
	}

	app := cli.App{
		Name:    "Dines",
		Usage:   `A disassembler for customed 8-bit microprocessor, "MOS Technology 6502" in Nintendo Entertainment System written in Golang.`,
		Version: version,
		Action:  run,
		Flags:   flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	romFile := c.String("rom")
	output := c.String("output")
	color := c.Bool("color")
	max := c.Int("max")

	data, err := ioutil.ReadFile(romFile)
	if err != nil {
		return err
	}

	result, err := disasm.Disassemble(data)
	if err != nil {
		return err
	}

	disasm.Dump(result, disasm.DumpMethod(output), color, max)

	return nil
}
