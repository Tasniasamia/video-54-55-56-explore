package cmd

import (
	"mains/config"
	"mains/rest/handlers/product"
	"mains/rest/handlers/user"
	"mains/rest"
	"mains/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig()

	middlewares := middleware.NewMiddleware(cnf);
	
	product :=product.NewHandler(middlewares);
	user :=user.NewHandler(middlewares);

  // server
   createServer :=rest.NewServer(cnf,product,user,middlewares);
   createServer.Start();
}
