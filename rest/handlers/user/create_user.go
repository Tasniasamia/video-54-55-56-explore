package user

import (
	"mains/util"
	"encoding/json"
	"net/http"
	"mains/repo"

)

type reqUser struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler)CreateUser(w http.ResponseWriter, r *http.Request){

	var newUser reqUser;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	insertUser:=h.userRepo.Resister(&repo.User{
		Id :len(h.userRepo.List())+1,
		Name: newUser.Name,
		Email:newUser.Email,
		Password: newUser.Password,
	});
	util.SendResponse(w, insertUser, http.StatusCreated)

}