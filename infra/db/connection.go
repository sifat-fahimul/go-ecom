package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	return "user=postgres password=123456789 host=localhost port=5432 dbname=ecommerce sslmode=disabled"
}

func NewConnection() (*sqlx.DB, error) {

	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil

}
