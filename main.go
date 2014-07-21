package main

import (
	"bufio"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bbq"
	app.Usage = "Tool to help you BBQ"
	app.Commands = []cli.Command{
		{
			Name:      "item",
			ShortName: "i",
			Usage:     "items needed for bbq",
			Action: func(c *cli.Context) {
				file, err := os.Open("items.txt")
				if err != nil {
					log.Fatal(err)
				}
        var reader *bufio.Reader
        var line []byte
				reader = bufio.NewReader(file)
				for {
					if err == io.EOF {
						return
					}
					line, err = reader.ReadBytes('\n')
					print(string(line))
				}
			},
		},
	}

	app.Run(os.Args)
}
