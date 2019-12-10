package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag(
		Name: "name, n",
		Value: "World",
		Usage: "Who to say hello to.",

	)
}
