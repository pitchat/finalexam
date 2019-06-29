package database

import (
	"database/sql"
)

// DataLayer is interface
type DataLayer interface {
	Insert(conn *sql.DB) (DataLayer, error)
	Update(conn *sql.DB) error
	Delete(conn *sql.DB) error
	GetByKey(conn *sql.DB) (DataLayer, error)
	GetAll(conn *sql.DB) ([]DataLayer, error)
}

// GetByKey table record by key
func GetByKey(d DataLayer) (DataLayer, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	o, err := d.GetByKey(conn)
	return o, err
}

// GetAll table records
func GetAll(d DataLayer) ([]DataLayer, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	o, err := d.GetAll(conn)
	return o, err
}

// Insert database
func Insert(d DataLayer) (DataLayer, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	o, err := d.Insert(conn)
	return o, err
}

// Update database
func Update(d DataLayer) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = d.Update(conn)
	return err
}

// Delete database
func Delete(d DataLayer) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = d.Delete(conn)
	return err
}

//IConv is convert struct to interface
func IConv(d DataLayer) DataLayer {
	return d
}
