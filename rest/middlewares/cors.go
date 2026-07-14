package middleware

import (
	"log"
	"net/http"
)

func (m *Middlewares) Cors(mux http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
	log.Println("Method is ",r.Method);
	//handle cors error
	w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
	w.Header().Set("Content-Type","application/json");
	
	 mux.ServeHTTP(w,r);
	}
	return http.HandlerFunc(handler);
}
