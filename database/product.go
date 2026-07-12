package database;

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

var users [] User;

func List() []User {
return users;
}

func  Store(u User) User {
	if(u.Id != 0) {
		return u;
	}
	u.Id=len(users)+1;
	users = append(users, u)
	return u;
}

func Find(id int) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil // Return zero value if not found
}

func Update(id int, updatedUser User) User {
	for i, user := range users {
		if user.Id == id {
			users[i] = updatedUser
			return updatedUser
		}
	}
	return User{} // Return zero value if not found
}

func Delete(id int) {
	var temp [] User;
	for _, user := range users {
		if user.Id != id {

			temp=append(temp, user)
			
		}
	}
	users = temp;
}


func init(){
	u1 := User{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := User{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,

		Dept: "Mathematics",
	}
	u3 := User{
		Id:   3,
		Name: "Bob Johnson",
		Roll: 103,
		Age:  24,
		Dept: "Physics",
	}
	users = append(users, u1, u2, u3)
}