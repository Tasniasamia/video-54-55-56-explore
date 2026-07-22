package product;
import (
	"mains/util"
	"net/http"
	

)


func(h *Handler)GetProducts(w http.ResponseWriter, r *http.Request) {

	users:=h.productRepo.List();
	util.SendResponse(w, users, http.StatusOK)
}