package rest

import (
	"fmt"
	"mains/config"
	"mains/rest/handlers/product"
	"mains/rest/handlers/user"
	"mains/rest/middlewares"
	"net/http"
)


type Server struct{
	cnf     config.Config
	product *product.Handler;
    user *user.Handler;
}

func NewServer(cnf config.Config,product *product.Handler,user *user.Handler)*Server{
	return &Server{
		cnf:cnf,
		product:product,
		user:user,
	}
}



func (server *Server) Start() {

   mux := http.NewServeMux()
	
   managerStruct:=middleware.NewManager();
   managerStruct.Use(middleware.Logger,middleware.Hudai,middleware.CorsPreflight);


server.product.ResisterRoutes(mux, managerStruct);
server.user.ResisterRoutes(mux, managerStruct);

wrapMux:=managerStruct.WrapMux(mux);


	fmt.Println("Server is running on : ", server.cnf.HttpPort);  
    address :=":"+fmt.Sprint(server.cnf.HttpPort)	
	err := http.ListenAndServe(address, wrapMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}