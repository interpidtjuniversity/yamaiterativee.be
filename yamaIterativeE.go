package main

import (
	"github.com/urfave/cli"
	_ "github.com/urfave/cli"
	"log"
	_ "log"
	"os"
	_ "os"
	"yama.io/yamaIterativeE/internal/cmd"
	_ "yama.io/yamaIterativeE/internal/cmd"
	"yama.io/yamaIterativeE/internal/conf"
	_ "yama.io/yamaIterativeE/internal/conf"
)

func main() {
	app := cli.NewApp()
	app.Name = "YamaIterativeE"
	app.Usage = "A painless self-hosted CI/CD service"
	app.Version = conf.App.Version
	app.Commands = []cli.Command{
		cmd.Web,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
