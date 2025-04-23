package database

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) UpdateAvatarURL(ethAddress, avatarURL string) (*entities.User, error) {
	query := `UPDATE users SET avatar_url = $1 WHERE eth_address = $2`
	_, err := u.db.Exec(query, avatarURL, ethAddress)
	if err != nil {
		return nil, fmt.Errorf("error updating avatar url: %w", err)
	}

	var nft entities.User
	query = "SELECT * FROM users WHERE eth_address = $1"

	err = u.db.Get(&nft, query, ethAddress)
	if err != nil {
		log.Printf("Db error: %v", err)
		return nil, fmt.Errorf("error updating avatar url: %w", err)
	}

	return &nft, nil
}

func (u *UserRepository) CheckUserExists(address string) (*bool, error) {
	var isExists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE eth_address = $1)"

	err := u.db.Get(&isExists, query, &address)

	if err != nil {
		return nil, fmt.Errorf("error checking if user exists: %w", err)
	}

	return &isExists, nil
}

func (u *UserRepository) GetUserByAddress(address string) (*entities.User, error) {
	var user entities.User
	query := "SELECT * FROM users WHERE eth_address = $1"
	err := u.db.Get(&user, query, address)

	if err != nil {
		return nil, fmt.Errorf("error getting user by address: %w", err)
	}

	return &user, nil
}

func (u *UserRepository) CreateUser(user *entities.User) error {
	query := "INSERT INTO users (eth_address) VALUES (:eth_address)"
	_, err := u.db.NamedExec(query, &user)

	if err != nil {
		log.Printf("DB error: %v", err)
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserRepository) GetAllUsers() (*[]entities.User, error) {
	var users []entities.User
	query := "SELECT * FROM users"
	err := u.db.Select(&users, query)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	return &users, nil
}

func (u *UserRepository) GetUserById(id string) (*entities.User, error) {
	var user entities.User
	query := "SELECT * FROM users WHERE id = ?"
	err := u.db.Get(&user, query, id)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return &user, nil
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{db: db}
}
