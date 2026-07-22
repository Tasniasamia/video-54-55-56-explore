package product

import (
	"encoding/json"
	"log"
	"mains/repo"
	"mains/util"
	"net/http"
)

type requestProduct struct {
	Name string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Category string `json:"category" db:"category"`
}

func(h *Handler)CreateProduct(w http.ResponseWriter, r *http.Request) {
   
	var newProduct requestProduct;
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	
	insertProduct,err:=h.productRepo.Store(repo.Product{
        Name:newProduct.Name,
		Price: newProduct.Price,
		Quantity: newProduct.Quantity,
		Category: newProduct.Category,

	});

	if(err != nil){
		log.Println(err);
	http.Error(w, err.Error(), http.StatusBadRequest)

	}

	log.Println("Inserted product:", insertProduct);
	util.SendResponse(w, insertProduct, http.StatusCreated)
}