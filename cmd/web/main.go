package main

import (
	"eon/kata/mike/pkg/domain/user"
	"eon/kata/mike/pkg/kernel"
)

func main() {
	app := kernel.Boot()

	user.InitDomain(app)

	app.Run()

	app.WaitForShutdown()
}
