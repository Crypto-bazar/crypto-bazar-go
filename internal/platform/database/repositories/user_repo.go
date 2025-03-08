package repositories

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// GetUserByAddress implements domain.UserRepository.
func (u *UserRepository) GetUserByAddress(address string) (*models.User, error) {
	var user models.User
	fmt.Println("Address to query:", address)
	query := "SELECT * FROM users WHERE eth_address = $1"
	err := u.db.Get(&user, query, address)

	if err != nil {
		return nil, fmt.Errorf("error getting user by address: %w", err)
	}

	return &user, nil
}

// CreateUser implements domain.UserRepository.
func (u *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (eth_address) VALUES (:eth_address)"
	_, err := u.db.NamedExec(query, &user)

	if err != nil {
		log.Printf("DB error: %v", err)
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
		log.Printf("DB error: %v", err)
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
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return &user, nil
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepository{db: db}
}
