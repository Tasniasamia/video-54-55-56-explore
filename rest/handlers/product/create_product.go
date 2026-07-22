package product;
import (
	"mains/util"
	"net/http"
	"encoding/json"
	"mains/repo"
	

)

type requestProduct struct {
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

func(h *Handler)CreateProduct(w http.ResponseWriter, r *http.Request) {
   
	var newUser requestProduct;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	
	insertUser:=h.productRepo.Store(repo.Product{
		Id:len(h.productRepo.List())+1,
		Name:newUser.Name,
		Roll: newUser.Roll,
		Age: newUser.Age,
		Dept: newUser.Dept,

	});
	util.SendResponse(w, insertUser, http.StatusCreated)
}