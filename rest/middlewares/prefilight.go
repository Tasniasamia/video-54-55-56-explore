package middleware

import (
	"log"
	"net/http"
)

func Preflight(mux http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
	
		if r.Method == http.MethodOptions {
		log.Print("options method")
		w.WriteHeader(http.StatusOK)
		return
	}
	 mux.ServeHTTP(w,r);
	}
	return http.HandlerFunc(handler);
}
