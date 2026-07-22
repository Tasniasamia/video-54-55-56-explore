package product

import  (
	"mains/rest/middlewares"
	"mains/repo"
)

type Handler struct {
middleware *middleware.Middlewares;
productRepo repo.ProductRepo;
}

func NewHandler(m *middleware.Middlewares,p repo.ProductRepo) *Handler {
	return &Handler{
		middleware: m,
		productRepo: p,
	}
}
