package product;
import (
	"mains/util"
	"mains/database"
	"net/http"
	

)


func(h *Handler)GetProducts(w http.ResponseWriter, r *http.Request) {

	users:=database.List();
	util.SendResponse(w, users, http.StatusOK)
}