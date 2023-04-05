package repository

import "github.com/jmoiron/sqlx"

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{db: db}
}

func (r CustomerRepositoryDB) GetAll() ([]Customer, error) {
	customer := []Customer{}
	query := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customer"
	err := r.db.Select(&customer, query)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r CustomerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customer WHERE customer_id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
