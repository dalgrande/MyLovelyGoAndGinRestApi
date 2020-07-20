package controllers //controllers/Product.go

import (
	config "api/config"
	"api/models"
	"fmt"
)

//GetAllProducts Fetch all product data
func GetAllProducts(product *[]models.Product) (err error) {
	if err = config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateProducts ... Insert New data
func CreateProducts(product *models.Product) (err error) {
	if err = config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductsByID ... Fetch only one product by Id
func GetProductsByID(product *models.Product, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProducts ... Update product
func UpdateProducts(product *models.Product, id string) (err error) {
	fmt.Println(product)
	config.DB.Save(product)
	return nil
}

//DeleteProducts ... Delete product
func DeleteProducts(product *models.Product, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(product)
	return nil
}
