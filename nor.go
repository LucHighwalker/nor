package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/initializer"
)

var binDir, _ = os.Executable()
var norDir = strings.TrimRight(binDir, "/nor") + "/nor"
var workDir, _ = os.Getwd()

var nor = cli.NewApp()

func clearTemp() {
	tempPath := path.Join(norDir, "__temp__")

	os.RemoveAll(tempPath)
	helper.EnsureDirExists(tempPath)
}

func info() {
	nor.Name = "Node on Rails"
	nor.Usage = "Like Ruby on Rails but NodeJS"
	author := cli.Author{Name: "Luc Highwalker", Email: "email@luc.gg"}
	nor.Authors = []*cli.Author{&author}
	nor.Version = "0.0.1"
}

func commands() {
	nor.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize a new NoR project.",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "defPort",
					Value: 4200,
					Usage: "Default port for server to run on.",
				},
				&cli.IntFlag{
					Name:  "dbPort",
					Value: 27017,
					Usage: "Default port mongo runs on.",
				},
			},
			Action: func(c *cli.Context) error {
				initializer.Initialize(norDir, workDir, c)
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add an existing module.",
			Action: func(c *cli.Context) error {
				// add module
				return nil
			},
		},
		{
			Name:    "controller",
			Aliases: []string{"c"},
			Usage:   "Create a new controller.",
			Action: func(c *cli.Context) error {
				// generate controller
				return nil
			},
		},
		{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "Generate a model.",
			Action: func(c *cli.Context) error {
				// Generate model
				return nil
			},
		},
		{
			Name:    "struct",
			Aliases: []string{"s"},
			Usage:   "Generate a structure (model/controller).",
			Action: func(c *cli.Context) error {
				// Generate a model struct
				return nil
			},
		},
	}
}

func main() {
	clearTemp()
	info()
	commands()
	err := nor.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
