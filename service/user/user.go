package user

import (
	"database/sql"
	"fmt"

	"github.com/pts/mdes/types"
)

type UserService struct {
	db *sql.DB
}

type MockUserService struct{}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (service *UserService) GetUserByEmail(email string) (*types.User, error) {
	rows, err := service.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found ..")
	}

	return u, nil
}

func (service *UserService) GetUserById(id int) (*types.User, error) {
	rows, err := service.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found ..")
	}

	return u, nil
}

func (service *UserService) CreateUser(user types.User) error {
	_, err := service.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Mocking UserService Methods

func (m *MockUserService) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User not found !")
}

func (m *MockUserService) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *MockUserService) CreateUser(types.User) error {
	return nil
}
