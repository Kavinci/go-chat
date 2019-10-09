package main

import (
	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "build"
	app.BuildCommand = ""
  	app.WindowsFlags = "-ldflags -H=windowsgui"
  	app.Action = func(c *cli.Context) error {
    println("Greetings")
    return nil
  }

  app.Run(os.Args)
}