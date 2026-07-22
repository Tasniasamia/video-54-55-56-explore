package repo;

type ProductRepo interface{
	Store(p Product)(*Product);
	List()[]*Product;
    Find(id int)*Product;
	Update(id int,updateUser *Product)(*Product)
	Delete(id int)
}



type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

type productRepo struct{
	productList []*Product;
}

func NewProductRepo()ProductRepo{
 repo :=&productRepo{

 };
 GenerateProduct(repo);
 return repo;
}

func (r *productRepo)List() []*Product {
return r.productList;
}

func(r *productRepo) Store(u Product) *Product {
	if(u.Id != 0) {
		return &u;
	}
	u.Id=len(r.productList)+1;
	r.productList = append(r.productList, &u)
	return &u;
}

func (r *productRepo)Find(id int) *Product {
	for _, user := range r.productList {
		if user.Id == id {
			return user
		}
	}
	return nil 
}

func (r *productRepo)Update(id int, updatedUser *Product) *Product {
	for i, user := range r.productList {
		if user.Id == id {
			r.productList[i] = updatedUser
			return updatedUser
		}
	}
	return &Product{} // Return zero value if not found
}

func (r *productRepo)Delete(id int){
	var temp [] *Product;
	for _, user := range r.productList{
		if user.Id != id {

			temp=append(temp, user)
			
		}
	}
	r.productList = temp;
	

}


func GenerateProduct(r *productRepo){
	u1 := &Product{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := &Product{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,

		Dept: "Mathematics",
	}
	u3 := &Product{
		Id:   3,
		Name: "Bob Johnson",
		Roll: 103,
		Age:  24,
		Dept: "Physics",
	}
	r.productList = append(r.productList, u1, u2, u3)
}