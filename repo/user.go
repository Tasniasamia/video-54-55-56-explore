package repo;

type User struct{
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRepo interface{
	Resister(a *User)*User;
	Login(email,password string) *User
	List() []*User
}


type userRepo struct{
	userList[]*User;
}

func NewUserRepo()UserRepo{
	return &userRepo{};
}


func(u *userRepo) Resister(a *User) *User {
	
	a.Id = len(u.userList) + 1;
	u.userList = append(u.userList, a);
	return a;
}	

func (u *userRepo) Login(email, password string) *User {
	for _, user := range u.userList {
		if user.Email == email && user.Password == password {
			return user
		}
	}
	return nil
}

func (u *userRepo)List() []*User {
return u.userList;
}