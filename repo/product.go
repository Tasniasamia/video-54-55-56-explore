package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)


type ProductRepo interface{
	Store(u Product) (*Product,error);
	List() ([]*Product,error);
    // Find(id int)*Product;
	// Update(id int,updateUser *Product)(*Product)
	// Delete(id int)
}



type Product struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Category string `json:"category" db:"category"`
}

type productRepo struct{
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB)ProductRepo{
 repo :=&productRepo{db:db };
//  GenerateProduct(repo);
 return repo;
}

func (r *productRepo)List() ([]*Product,error) {
	query := `SELECT * FROM product`
    var productList []*Product;
	err:=r.db.Select(&productList,query);
	if(err != nil){
		
	return nil,err
	}
    
	
return productList,nil;
}

func(r *productRepo) Store(u Product) (*Product,error) {
	if(u.Id != 0) {
		return &u,nil;
	}
	
	query := `
    INSERT INTO product (name, price, quantity, category)
    VALUES (:name, :price, :quantity, :category)
    RETURNING *`
    var product *Product= &Product{}
	row,err:=r.db.NamedQuery(query, u)
	if(err != nil){
      return nil,err
	}
    
	
      
if row.Next() {
   row.StructScan(product)
}
	
    log.Println("product ",product);
	return product,nil;
	


}

// func (r *productRepo)Find(id int) *Product {
// 	for _, user := range r.productList {
// 		if user.Id == id {
// 			return user
// 		}
// 	}
// 	return nil 
// }

// func (r *productRepo)Update(id int, updatedUser *Product) *Product {
// 	for i, user := range r.productList {
// 		if user.Id == id {
// 			r.productList[i] = updatedUser
// 			return updatedUser
// 		}
// 	}
// 	return &Product{} // Return zero value if not found
// }

// func (r *productRepo)Delete(id int){
// 	var temp [] *Product;
// 	for _, user := range r.productList{
// 		if user.Id != id {

// 			temp=append(temp, user)
			
// 		}
// 	}
// 	r.productList = temp;
	

// }


// func GenerateProduct(r *productRepo){
// 	u1 := &Product{
// 		Id:   1,
// 		Name: "John Doe",
// 		Roll: 101,
// 		Age:  25,
// 		Dept: "Computer Science",
// 	}
// 	u2 := &Product{
// 		Id:   2,
// 		Name: "Jane Smith",
// 		Roll: 102,
// 		Age:  23,

// 		Dept: "Mathematics",
// 	}
// 	u3 := &Product{
// 		Id:   3,
// 		Name: "Bob Johnson",
// 		Roll: 103,
// 		Age:  24,
// 		Dept: "Physics",
// 	}
// 	r.productList = append(r.productList, u1, u2, u3)
// }