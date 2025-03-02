package repositories

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// CreateUser implements domain.UserRepository.
func (u *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INFO users (eth_address) VALUES (:eth_address)"
	_, err := u.db.NamedExec(query, &user)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

// GetAllUsers implements domain.UserRepository.
func (u *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users"
	err := u.db.Get(&users, query)

	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	return &users, nil
}

// GetUserById implements domain.UserRepository.
func (u *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id = ?"
	err := u.db.Get(&user, query, id)

	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return &user, nil
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepository{db: db}
}
