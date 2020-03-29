package main

import (
	"goss/Colors"
	"goss/Commands"
	"os"
	"github.com/urfave/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "goss"
	app.Usage = "set up oss contributing"
	app.Version = "0.0.1"

	app.Commands = []*cli.Command{
		{
			Name:    "create",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				Commands.Create(c.Args().Get(0),c.Args().Get(1))
				fmt.Println( Colors.Green("ALL success!!") )
				return nil
			},
		},
		{
			Name:    "issue",
			Aliases: []string{"a"},
			Usage:   "issueが確認できるよ。--でoptionを追加できるよ・　--label でラベルを指定できるよ　--sort IDの順番を指定すると降順昇順で表示できるよ　--length 表示数を確認できるよ",
			Action: func(c *cli.Context) error {
				fmt.Println("issueだよ")
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"a"},
			Usage:   "issueブランチの詳細が確認できるよ。upsreamがにそのブランチがあるあも確認できます",
			Action: func(c *cli.Context) error {
				fmt.Println("showだよ")
				return nil
			},
		},
	}
	app.Run(os.Args)
}