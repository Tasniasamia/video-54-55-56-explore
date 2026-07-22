package product;
import (
	"mains/util"
	"net/http"
	

)


func(h *Handler)GetProducts(w http.ResponseWriter, r *http.Request) {

	users,err:=h.productRepo.List();
	if(err != nil){
			http.Error(w, err.Error(), http.StatusBadRequest)

	}
	util.SendResponse(w, users, http.StatusOK)
}