package infrastructureU

import (
	"ArquitecturaExagonal/src/users/domainU/userEntity"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Save(user *userEntity.User) error {
	query := "INSERT INTO users (name, phone_number) VALUES (?, ?)"
	fmt.Printf("Guardando en DB: Name=%s, Phone=%s\n", user.Name, user.Phone)
	_, err := repo.db.Exec(query, user.Name, user.Phone)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

func (repo *UserRepository) GetAll() ([]*userEntity.User, error) {
	query := "SELECT * FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	defer rows.Close()
	var users []*userEntity.User
	for rows.Next() {
		var user userEntity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone); err != nil {
			return nil, fmt.Errorf("Error: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) Update(id int32, user *userEntity.User) error {
	query := "UPDATE users SET name = ?, phone_number = ? WHERE id = ?"
	_, err := repo.db.Exec(query, user.Name, user.Phone, id)
	return err
}

func (repo *UserRepository) GetByID(id int32) (*userEntity.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := repo.db.QueryRow(query, id)
	var user userEntity.User
	err := row.Scan(&user.ID, &user.Name, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	return &user, nil
}

func (repo *UserRepository) Delete(id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}
