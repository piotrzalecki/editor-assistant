package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

var model string

var appConfig = &cli.App{
	Name:  "editor-assistant",
	Version: "v0.0.1",
	Usage: "This is editor assistant, that read from clipboard, edit text and put it back to clipboard.",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("No arguments provided")
		return errors.New("no arguments provided, use 'help' to see available options")
	},
	Commands: []*cli.Command{
		{
			Name:    "ollama",
			Aliases: []string{"o"},
			Usage:   "use local Ollama model",
			Action: func(cCtx *cli.Context) error {
					return watchInput()

			},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "model",
				Aliases: []string{"m"},
				Value:   "editor-assistant:gemma2",
				Usage:   "Local Ollama model to use.",
				EnvVars: []string{"EDITOR-ASSISTANT_OLLAMA_MODEL"},
				Destination: &model,
			},
		},
	},
	{
			Name:    "gpt",
			Aliases: []string{"g"},
			Usage:   "use ChatGPT model",
			Action: func(cCtx *cli.Context) error {
				return nil

			},
	},
	},
}