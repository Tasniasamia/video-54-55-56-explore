package user;
import (
	"net/http"
    "mains/rest/middlewares"
)

func (h *Handler)ResisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /register", manager.With(http.HandlerFunc(h.CreateUser)));
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.Login)));
}








