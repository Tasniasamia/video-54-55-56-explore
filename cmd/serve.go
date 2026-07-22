package cmd;

import (
    "mains/config"
	"mains/rest"
	"mains/rest/handlers/product"
	"mains/rest/handlers/user"
	"mains/rest/middlewares"
	"mains/repo"

)

func Serve() {

	cnf := config.GetConfig()
    middlewares := middleware.NewMiddleware(cnf);
	
	productRepo:=repo.NewProductRepo();
	product :=product.NewHandler(middlewares,productRepo);

	userRepo:=repo.NewUserRepo();
	user :=user.NewHandler(middlewares,userRepo);

  // server
   createServer :=rest.NewServer(cnf,product,user,middlewares);
   createServer.Start();
}
