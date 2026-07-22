package product;

import (
	"net/http"
	"mains/rest/middlewares"
)

func (h *Handler)ResisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.HandleFunc("GET /", (http.HandlerFunc(handler.GetProducts)).ServeHTTP);
	// mux.HandleFunc("POST /product", (http.HandlerFunc(handler.CreateProduct)).ServeHTTP)
	// mux.HandleFunc("GET /product/{id}", (http.HandlerFunc(handler.GetSingleProduct)).ServeHTTP)

	
	mux.Handle("GET /",manager.With(http.HandlerFunc(h.GetProducts),h.middleware.Logger,h.middleware.Hudai));
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct),h.middleware.Logger,h.middleware.Hudai));
	// mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetSingleProduct),h.middleware.Logger,h.middleware.Hudai));
	// mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct),h.middleware.Logger,h.middleware.Hudai));
	// mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct),h.middleware.Logger,h.middleware.Hudai));
	
	
}






// func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
// 	if use121 {
// 		return mux.mux121.findHandler(r)
// 	}
// 	h, p, _, _ := mux.findHandler(r)
// 	return h, p
// }



