package cmd

import (
	"mains/config"
	"mains/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)
}
