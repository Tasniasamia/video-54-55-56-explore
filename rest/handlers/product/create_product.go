package product;
import (
	"mains/util"
	"mains/database"
	"net/http"
	"encoding/json"
	

)


func  CreateProduct(w http.ResponseWriter, r *http.Request) {
   
	var newUser database.User;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	
	insertUser:=database.Store(newUser)
	util.SendResponse(w, insertUser, http.StatusCreated)
}