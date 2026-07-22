package product;
import (
	"mains/util"
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"mains/repo"
	
)

type requestUpdateProduct struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}
func(h *Handler)UpdateProduct(w http.ResponseWriter, r *http.Request) {

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

   var findUser requestUpdateProduct;
	err = json.NewDecoder(r.Body).Decode(&findUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

  fmt.Println("findUser:", findUser)
   findUser.Id=productIdInt;
    getUserById:=h.productRepo.Update(productIdInt,&repo.Product{
		Id:productIdInt,
		Name:findUser.Name,
		Roll: findUser.Roll,
		Dept: findUser.Dept,
		Age: findUser.Age,
	})
	fmt.Println("getUserById:", getUserById)
	if(getUserById.Id == productIdInt) {
		util.SendResponse(w, getUserById, http.StatusOK)
		return
	}
	http.Error(w, "Product not updated", http.StatusNotFound)	
}
