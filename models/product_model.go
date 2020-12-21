package models

import (
	"context"
	"database/sql"
	"mockdb/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func CreateProductModel(db *sql.DB) *ProductModel {
	productModel := &ProductModel{Db: db}
	return productModel
}

//FindAll
func (productModel ProductModel) FindAll(ctx context.Context) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

//SearchID
func (productModel ProductModel) Search(ctx context.Context, keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}

		}
		return products, nil
	}
}
