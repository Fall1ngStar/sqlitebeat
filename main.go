package main

import (
	"os"

	"github.com/Fall1ngStar/sqlitebeat/cmd"

	_ "github.com/Fall1ngStar/sqlitebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
