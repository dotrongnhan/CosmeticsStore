package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"fmt"
)

func GetProducts(sortType string, category string, limit int, offset int) (products []models.Product, countProducts int) {
	offsetValue := limit * (offset - 1)
	query, err := database.DB.Query("")
	count, err := database.DB.Query("")
	fmt.Println(sortType,category,offset)
	if category == "Default" && sortType == "ASC" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.price ASC LIMIT ? OFFSET ?", limit, offsetValue)
		count, _ = database.DB.Query("SELECT COUNT(*) FROM products")
	} else if sortType == "DESC" && category == "Default" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.price DESC LIMIT ? OFFSET ?", limit, offsetValue)
		count, _ = database.DB.Query("SELECT COUNT(*) FROM products")
	} else if category != "Default" && sortType == "ASC" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products as P JOIN categories as C ON C.id = P.category_id JOIN brands as B ON B.id = P.brand_id where C.category_name = ? ORDER BY P.price ASC LIMIT ? OFFSET ?", category, limit, offsetValue)
		count, _ = database.DB.Query("SELECT count(*) FROM products as P JOIN categories as C ON C.id = P.category_id JOIN brands as B ON B.id = P.brand_id where C.category_name = ?", category)
	} else if category != "Default" && sortType == "DESC" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products as P JOIN categories as C ON C.id = P.category_id JOIN brands as B ON B.id = P.brand_id where C.category_name = ? ORDER BY P.price DESC LIMIT ? OFFSET ?", category, limit, offsetValue)
		count, _ = database.DB.Query("SELECT count(*) FROM products as P JOIN categories as C ON C.id = P.category_id JOIN brands as B ON B.id = P.brand_id where C.category_name = ?", category)
	} else {
		fmt.Println(1)
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id LIMIT ? OFFSET ?", limit, offsetValue)
		count, _ = database.DB.Query("SELECT COUNT(*) FROM products")
	}

	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
	}

	for count.Next() {
		count.Scan(&countProducts)
	}

	return products, countProducts
}

func SearchProducts(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE P.product_name LIKE ?", "%"+name+"%")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products
}

func GetProductByCategoryName(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE C.category_name = ?", name)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
	}
	return products
}

func GetProductByBrandName(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE C.category_name = ?", name)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
	}
	return products
}

func GetProduct(id int) (product models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE P.id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
	}
	return product
}

func CreateProduct(product models.Product) error {
	stmt, err := database.DB.Prepare("INSERT INTO products(product_name, description, price, image, is_hot, category_id, brand_id, number_available) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(product.ProductName, product.Description, product.Price, product.Image, product.IsHot, product.CategoryId, product.BrandId, product.NumberAvailable)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateProduct(product models.Product, id int) error {
	stmt, err := database.DB.Prepare("UPDATE products SET product_name=?,description=?,price=?,image=?,is_hot=?,category_id=?,brand_id=?,number_available=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.Description, product.Price, product.Image, product.IsHot, product.CategoryId, product.BrandId, product.NumberAvailable, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteProduct(id int) error {
	query, err := database.DB.Query("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
