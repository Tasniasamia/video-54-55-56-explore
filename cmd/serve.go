package cmd

import (
	"mains/config"
	"mains/rest/handlers/product"
	"mains/rest/handlers/user"
	"mains/rest"
)

func Serve() {

	cnf := config.GetConfig()

    //call All handler Function
	product :=product.NewHandler();
	user :=user.NewHandler();

  // New Server call
   createServer :=rest.NewServer(cnf,product,user);
   createServer.Start();
}
