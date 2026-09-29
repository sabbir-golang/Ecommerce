package database

import "fmt"

type Product struct {
	ID          int    `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

var productList []Product

func Store(p Product) {
	productList = append(productList, p)
}
func List() []Product {
	return productList
}
func Get(productId int) *Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}
func Update(UpdateProduct Product) {
	for i, product := range productList {
		if product.ID == UpdateProduct.ID {
			productList[i].Title = UpdateProduct.Title
			productList[i].Description = UpdateProduct.Description
			productList[i].Price = UpdateProduct.Price
			fmt.Println(productList[i])
			return
		}
	}
}
func Delete(productid int) {
	var tmp []Product
	for _, product := range productList {
		if product.ID != productid {
			tmp = append(tmp, product)
		}
	}
	productList = tmp
}
