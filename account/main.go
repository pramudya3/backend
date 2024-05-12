package main

import (
	"github.com/pramudya3/backend/payment/cmd"
)

func main() {
	app := cmd.New()

	cmd.InitSupertokens(app.Config)

	app.Start()
}
