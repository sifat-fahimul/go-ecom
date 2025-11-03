package repo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	ImgUrl      string    `json:"imageUrl" db:"image_url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Delete(id int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}

}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `INSERT INTO products (title, description, price, image_url) VALUES (:title, :description, :price, :image_url) RETURNING id`
	row, err := r.db.NamedQuery(query, &p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if row.Next() {
		row.Scan(&p.ID)
	}
	return &p, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	var product Product
	err := r.db.Get(&product, `SELECT * FROM products WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product
	err := r.db.Select(&products, `SELECT * FROM products`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return products, nil
}
func (r *productRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	fmt.Println(err)
	return err
}

func (r *productRepo) Update(product Product) (*Product, error) {
	query := `UPDATE products SET title = :title, description = :description, price = :price, image_url = :image_url, updated_at = NOW() WHERE id = :id`
	_, err := r.db.NamedExec(query, &product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}
