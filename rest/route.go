package rest;

import (
	"net/http"

	 "mains/rest/handlers/product"
	 "mains/rest/handlers/user"
	 "mains/rest/middlewares"

)

func initiateRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.HandleFunc("GET /", (http.HandlerFunc(handler.GetProducts)).ServeHTTP);
	// mux.HandleFunc("POST /product", (http.HandlerFunc(handler.CreateProduct)).ServeHTTP)
	// mux.HandleFunc("GET /product/{id}", (http.HandlerFunc(handler.GetSingleProduct)).ServeHTTP)

	
	mux.Handle("GET /",manager.With(http.HandlerFunc(product.GetProducts)));
	mux.Handle("POST /products", manager.With(http.HandlerFunc(product.CreateProduct)));
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(product.GetSingleProduct)));
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(product.UpdateProduct)));
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(product.DeleteProduct)));
	
	mux.Handle("POST /register", manager.With(http.HandlerFunc(user.CreateUser)));
	mux.Handle("POST /login", manager.With(http.HandlerFunc(user.Login)));
}






// func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
// 	if use121 {
// 		return mux.mux121.findHandler(r)
// 	}
// 	h, p, _, _ := mux.findHandler(r)
// 	return h, p
// }



