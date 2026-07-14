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
	middlewares *middleware.Middlewares;
}

func NewServer(cnf config.Config,product *product.Handler,user *user.Handler,middlewares *middleware.Middlewares)*Server{
	return &Server{
		cnf:cnf,
		product:product,
		user:user,
		middlewares:middlewares,
	}
}



func (server *Server) Start() {

   mux := http.NewServeMux()
	
   managerStruct:=middleware.NewManager();
   managerStruct.Use(server.middlewares.Logger,server.middlewares.Hudai,server.middlewares.CorsPreflight);

//routes
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