package user

import  "mains/rest/middlewares"
import "mains/repo"

type Handler struct {
	middleware *middleware.Middlewares;
	userRepo repo.UserRepo;

}

func NewHandler(m *middleware.Middlewares,u repo.UserRepo) *Handler {
	return &Handler{middleware: m,userRepo:u}
}
