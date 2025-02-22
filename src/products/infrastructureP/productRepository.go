package infrastructureP

import (
	"ArquitecturaExagonal/src/products/domainP/productEntity"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) Save(product *productEntity.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?,?)"
	_, err := repo.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

func (repo *ProductRepository) GetAll() ([]*productEntity.Product, error) {
	query := "SELECT * FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	defer rows.Close()

	var products []*productEntity.Product
	for rows.Next() {
		var product productEntity.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("Error: %w", err)
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repo *ProductRepository) Update(id int32, product *productEntity.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := repo.db.Exec(query, product.Name, product.Price, id)
	return err
}

func (repo *ProductRepository) Delete(id int32) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (pr *ProductRepository) GetByID(id int32) (*productEntity.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = ?"
	row := pr.db.QueryRow(query, id)

	var product productEntity.Product
	err := row.Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("producto con id %d no encontrado", id)
		}
		return nil, err
	}

	return &product, nil
}
