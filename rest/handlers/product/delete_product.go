package product;
import (
	"mains/database"
	"net/http"
	"strconv"

)

func(h *Handler)DeleteProduct(w http.ResponseWriter, r *http.Request) {

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
	database.Delete(productIdInt);
	http.Error(w, "Product deleted", http.StatusOK)	
}
