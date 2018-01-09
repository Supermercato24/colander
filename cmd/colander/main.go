// Package main implements methods to colander log aggregator.
//
// Main.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/supermercato24/colander/config"
	"github.com/urfave/cli"
)

const (
	ExIoerr = 74 // Exit Error
)

func init() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version number and exit",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%s %s\n", strings.Title(config.Name), config.Version)
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show this help message and exit",
	}
	cli.AppHelpTemplate = `usage: {{.HelpName}} {{if .VisibleFlags}}[optional options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}

{{.Name}} - {{.Usage}}{{if not .HideHelp}}{{ "\n" }}{{end}}
{{if .Commands}}positional arguments:
{{range .Commands}}{{if not .HideHelp}}  {{join .Names ", "}}{{ "\t" }}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}
{{if .VisibleFlags}}optional arguments:
  {{range .VisibleFlags}}{{.}}
  {{end}}
{{end}}`
}

func main() {
	var (
		dir     string
		pattern string
		remove  bool
		show    bool
	)

	app := cli.NewApp()
	app.Name = config.Name
	app.Version = config.Version
	app.Usage = "aggregate logs"
	//app.HideHelp = true
	//app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "d, dir",
			Usage:       "load logs from `DIR`",
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "p, pattern",
			Usage:       "load logs with this `PATTERN`",
			Destination: &pattern,
		},
		cli.BoolFlag{
			Name:        "r, remove",
			Usage:       "remove parsed logs",
			Destination: &remove,
		},
		cli.BoolFlag{
			Name:        "s, show",
			Usage:       "show output to screen",
			Destination: &show,
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c)

		fmt.Fprint(c.App.Writer, "command not found\n\n")
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		fmt.Fprint(c.App.Writer, "incorrect Usage\n\n")

		cli.ShowAppHelp(c)

		return nil
	}
	app.Action = func(c *cli.Context) error {

		if dir != "" {
			fileInfo, err := os.Stat(dir)
			if err != nil {
				return cli.NewExitError("input DIR doesn't not exist", ExIoerr)
			}
			mode := fileInfo.Mode()

			if exist := mode.IsDir(); !exist {
				return cli.NewExitError("input DIR is not a directory", ExIoerr)
			}
		} else if pattern == "" {
			return cli.NewExitError("input DIR is not a directory", ExIoerr)
		}

		colander(&colanderOptions{
			dir:     dir,
			pattern: pattern,
			remove:  remove,
			show:    show,
		})

		return nil
	}

	app.Run(os.Args)
}
