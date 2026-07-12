package middleware

import (
	"log"
	"net/http"
)

func CorsPreflight(mux http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
	w.Header().Set("Content-Type","application/json");
		if r.Method == http.MethodOptions {
		log.Print("options method")
		w.WriteHeader(http.StatusOK)
		return
	}
	 mux.ServeHTTP(w,r);
	}
	return http.HandlerFunc(handler);
}
