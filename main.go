package main

import (
	"github.com/ZeyuSi-2099/zema-cli/cmd"
	"github.com/ZeyuSi-2099/zema-cli/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
