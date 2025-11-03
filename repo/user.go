package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"isShopOwner" db:"isShopOwner"`
}

type UserRepo interface {
	Create(p User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Delete(id int) error
	// Update(p User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {

	query := `INSERT INTO users (firstName, lastName, email, password, isShopOwner) VALUES (:firstName, :lastName, :email, :password, :isShopOwner) RETURNING id`

	row, err := r.db.NamedQuery(query, &user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if row.Next() {
		row.Scan(&user.ID)
	}
	return &user, nil
}
func (r *userRepo) Find(email, password string) (*User, error) {
	var user User
	err := r.db.Get(&user, `SELECT * FROM users WHERE email = $1 AND password = $2 LIMIT 1`, email, password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

// func (r userRepo) Delete(id int) error
// func (r userRepo) Update(p User) (*User, error)
