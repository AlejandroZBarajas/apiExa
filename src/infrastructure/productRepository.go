package infrastructure

import (
	"ArquitecturaExagonal/src/domain/entities"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) Save(product *entities.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?,?)"
	_, err := repo.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

func (repo *ProductRepository) GetAll() ([]*entities.Product, error) {
	query := "SELECT * FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("Error: %w", err)
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repo *ProductRepository) Update(id int32, name string, price float32) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *ProductRepository) Delete(id int32) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}
