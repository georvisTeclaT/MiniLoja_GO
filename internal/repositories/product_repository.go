package repositories

import (
	"database/sql"
	"mini-loja/internal/dto/product"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) interfaces.IProductRepository {
	return productRepository{db: db}
}

func (p productRepository) GetAll() ([]product.ProductDto, error) {
	rows, err := p.db.Query("SELECT id, product_name, price FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsList []product.ProductDto
	for rows.Next() {
		var product product.ProductDto
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		productsList = append(productsList, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productsList, nil
}

func (p productRepository) GetByID(id int) (product.ProductDto, error) {
	rows, err := p.db.Query("SELECT id, product_name, price FROM product WHERE id = $1", id)
	if err != nil {
		return product.ProductDto{}, err
	}
	defer rows.Close()

	var product product.ProductDto
	if rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return product, err
		}
	}

	return product, sql.ErrNoRows
}

func (p productRepository) GetProductByID(id int) (models.Product, error) {
	rows, err := p.db.Query("SELECT id, product_name, price FROM product WHERE id = $1", id)
	if err != nil {
		return models.Product{}, err
	}
	defer rows.Close()

	var product models.Product
	if rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return product, err
		}
	}

	return product, sql.ErrNoRows
}

func (p productRepository) Create(prod models.Product) error {
	return p.db.QueryRow("INSERT INTO product(product_name, price) VALUES($1, $2) RETURNING id", prod.Name, prod.Price).Scan(&prod.Id)
}

func (p productRepository) Update(prod models.Product) error {
	return p.db.QueryRow("UPDATE product SET product_name=$1, price=$2 WHERE id=$3", prod.Name, prod.Price, prod.Id).Scan(&prod.Id)
}

func (p productRepository) Delete(id int) error {
	return p.db.QueryRow("DELETE FROM product WHERE id = $1", id).Scan(&id)
}
