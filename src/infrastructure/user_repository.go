package infrastructure

import (
	"ArquitecturaExagonal/src/domain/entities"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Save(user *entities.User) error {
	query := "INSERT INTO users (name, phone_number) VALUES (?, ?)"
	_, err := repo.db.Exec(query, user.Name, user.PhoneNumber)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

func (repo *UserRepository) GetAll() ([]*entities.User, error) {
	query := "SELECT * FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	defer rows.Close()
	var users []*entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.PhoneNumber); err != nil {
			return nil, fmt.Errorf("Error: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) Update(id int32, user *entities.User) error {
	query := "UPDATE users SET name = ?, phone_number = ? WHERE id = ?"
	_, err := repo.db.Exec(query, user.Name, user.PhoneNumber, id)
	return err
}

func (repo *UserRepository) Delete(id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *UserRepository) GetByID(id int32) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := repo.db.QueryRow(query, id)
	var user entities.User
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	return &user, nil
}
