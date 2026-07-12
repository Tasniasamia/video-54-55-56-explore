package middleware;
import (
	"log"
	"net/http"
);

func Hudai(next http.Handler) http.Handler {
	handler:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	   log.Println("Ami holam Hudai");
	   next.ServeHTTP(w, r)
	})
	return handler
}
