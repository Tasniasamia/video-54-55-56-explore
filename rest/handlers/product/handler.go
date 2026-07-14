package product

import  "mains/rest/middlewares"

type Handler struct {
middleware *middleware.Middlewares;

}

func NewHandler(m *middleware.Middlewares) *Handler {
	return &Handler{
		middleware: m,
	}
}
