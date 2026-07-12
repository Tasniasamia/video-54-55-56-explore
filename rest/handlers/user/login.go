package user

import (
	"mains/util"
	"encoding/json"
	"mains/database"
	"net/http"
	"log"
	"mains/config"
);



type loginUser struct{
	Email string `json:"email"`;
	Password string `json:"password"`;
}
func Login(w http.ResponseWriter, r *http.Request){

	cnf :=config.GetConfig();

	var newUser database.AuthUser;
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	findUser:=database.Login(newUser.Email, newUser.Password);
	if(findUser == nil) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	log.Println(findUser);
	token :=util.CreateJwtToken(cnf.JwtSecret, util.Payload{Sub: findUser.Id, Name: findUser.Name, Email: findUser.Email});
	log.Println(token);
	util.SendResponse(w, map[string]interface{}{"user": findUser, "token": token}, http.StatusOK)

}