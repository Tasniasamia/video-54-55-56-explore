package rest

import (
	"fmt"
	"mains/config"
	"mains/rest/middlewares"
	
	"net/http"
)




func  Start(cnf config.Config) {

   mux := http.NewServeMux()
	
   	managerStruct:=middleware.NewManager();
	managerStruct.Use(middleware.Logger,middleware.Hudai,middleware.CorsPreflight);


initiateRoutes(mux, managerStruct);

wrapMux:=managerStruct.WrapMux(mux);


	fmt.Println("Server is running on : ", cnf.HttpPort);  
    address :=":"+fmt.Sprint(cnf.HttpPort)	
	err := http.ListenAndServe(address, wrapMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}