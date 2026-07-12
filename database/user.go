package database;

type AuthUser struct{
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var authUsers []AuthUser;

func(a AuthUser) Resister() AuthUser {
	if(a.Id != 0) {
		return a
	}
	a.Id = len(authUsers) + 1;
	authUsers = append(authUsers, a);
	return a;
}	

func Login(email, password string) *AuthUser {
	for _, user := range authUsers {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}