package main

import (
	_ "github.com/nickmro/world-app/world/cmd"
	"github.com/nickmro/world-app/world/cmd"
	"os"
)

func main() {
	cmd.NewCli().Run(os.Args)
}
