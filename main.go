package main

import (
	"bufio"
	"fmt"
	"github.com/aybabtme/rgbterm"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"

	"github.com/urfave/cli"
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
				fmt.Println("Input your github ID:")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				username := scanner.Text()
				fmt.Println("Input your github password:")
				password, err := terminal.ReadPassword(int(syscall.Stdin))
				if err != nil {
					log.Fatal("なんか起きた")
				}
				var r, g, b uint8
				r, g, b = 252, 255, 43
				word := string(password)
				coloredWord := rgbterm.FgString(word, r, g, b)
				fmt.Println("Oh!", coloredWord, "hello!",username)
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