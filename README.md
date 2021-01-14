# Dines

![](https://img.shields.io/github/v/release/0n1shi/dines?sort=date&color=success) ![license](https://img.shields.io/badge/license-MIT-blue)

Disassembler for customed 8-bit microprocessor, "MOS Technology 6502" in Nintendo Entertainment System written in Golang.

## Usage

```
NAME:
   Dines - A disassembler for customed 8-bit microprocessor, "MOS Technology 6502" in Nintendo Entertainment System written in Golang.

USAGE:
   dines [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --rom value     A file path of NES ROM
   --output value  output format, "json" or "yaml", default is like a typical diassembler (default: normal)
   --color         color output (*only available without "output" option) (default: false)
   --max value     max  number of lines of output excluding header information (*only available without "output" option) (default: -1)
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```



## Libs

- https://github.com/urfave/cli
- https://github.com/stretchr/testify
- https://github.com/fatih/color
## Refs

- https://wiki.nesdev.com/w/index.php/INES
