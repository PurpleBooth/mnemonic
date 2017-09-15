// 	mnemonic - a terrible mnemonic generator
// 	Copyright (C) 2017 Billie Alice Thompson
//
// 	This program is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
// 	This program is distributed in the hope that it will be useful,
// 	but WITHOUT ANY WARRANTY; without even the implied warranty of
// 	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// 	GNU General Public License for more details.
//
// 	You should have received a copy of the GNU General Public License
// 	along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lloyd/wnram"
	"github.com/purplebooth/mnemonic/mnemonic"
	"github.com/urfave/cli"
)

const (
	// ErrorExitCodeWordNet is the exit code for a word net error
	ErrorExitCodeWordNet = 1 << iota
	// ErrorExitCodeTemplateParseError is the exit code for a template parse error
	ErrorExitCodeTemplateParseError
)

func main() {
	app := cli.NewApp()
	app.Version = "v0.1.0"
	app.Authors = []cli.Author{
		{
			Name:  "Billie Alice Thompson",
			Email: "billie@purplebooth.co.uk",
		},
	}
	app.Name = "mnemonic"
	app.Description = "Generate a mnemonic from a string of characters"
	app.Copyright = `
	mnemonic  Copyright (C) 2017  Billie Alice Thompson
	This program comes with ABSOLUTELY NO WARRANTY;	This is free software,
	and you are welcome to redistribute it under certain conditions; see
	LICENSE.md for additional details.
	`
	app.Commands = []cli.Command{
		{
			Name:        "generate",
			ArgsUsage:   "[PATH-TO-DICTIONARY] [LETTERS]",
			Usage:       "Generate a mnemonic from a string of characters (Default)",
			Description: "Generate a mnemonic from a string of characters",

			Action: func(c *cli.Context) error {
				dictDir := c.Args().Get(0)
				letters := strings.Split(strings.ToLower(c.Args().Get(1)), "")
				template := mnemonic.NewTemplate(letters)

				// This library is bad and prints to stdout.
				// We're swapping out stdout for a fake here then restoring it
				rescueStdout := os.Stdout
				_, w, _ := os.Pipe()
				os.Stdout = w

				wn, err := wnram.New(dictDir)

				w.Close()
				os.Stdout = rescueStdout

				if err != nil {
					log.Fatal(err)

					return cli.NewExitError(err.Error(), ErrorExitCodeWordNet)
				}

				generator := mnemonic.NewTemplateParser(
					mnemonic.NewWnramWordGenerator(wn, wnram.Adjective),
					mnemonic.NewWnramWordGenerator(wn, wnram.Noun),
					mnemonic.NewWnramWordGenerator(wn, wnram.Verb),
					mnemonic.NewWnramWordGenerator(wn, wnram.Adverb),
				)

				err = generator.Parse(template, letters, bufio.NewWriter(os.Stdout))

				fmt.Println()

				if err != nil {
					log.Fatal(err)
					return cli.NewExitError(err.Error(), ErrorExitCodeTemplateParseError)
				}

				return nil
			},
		},
	}

	app.Action = app.Commands[0].Action

	app.EnableBashCompletion = true
	app.Run(os.Args)
}
