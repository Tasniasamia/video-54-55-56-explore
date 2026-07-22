package product;
import (
	"mains/util"
	"net/http"
	"strconv"
	

)
func(h *Handler)GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	productId :=r.PathValue("id");
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if(productId == "") {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	

	getUserById := h.productRepo.Find(productIdInt);
	if(getUserById.Id == productIdInt) {
		util.SendResponse(w, getUserById, http.StatusOK)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)	
}