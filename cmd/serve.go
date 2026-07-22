package cmd

import (
	"mains/config"
	"mains/infra/db"
	"mains/repo"
	"mains/rest"
	"mains/rest/handlers/product"
	"mains/rest/handlers/user"
	"mains/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig();
	dbConnection:=db.NewConnection(*cnf.Dbconfig);
    middlewares := middleware.NewMiddleware(cnf);
	
	productRepo:=repo.NewProductRepo(dbConnection);
	product :=product.NewHandler(middlewares,productRepo);

	userRepo:=repo.NewUserRepo();
	user :=user.NewHandler(middlewares,userRepo);

  // server
   createServer :=rest.NewServer(cnf,product,user,middlewares);
   createServer.Start();
}
