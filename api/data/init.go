package data

import (
	"github.com/jmoiron/sqlx"
	"order-management-service/api/data/dao"
	"os"
	"time"
)

func Initialize() error {
	url := os.Getenv("SQL_URL")
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	db, err = sqlx.Open("mysql", url)
	if err != nil {
		return err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dao.OrdersDao.Connect()
	// Create the orders table if it doesn't already exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
        id VARCHAR(36) NOT NULL PRIMARY KEY,
        status VARCHAR(20) NOT NULL,
        items JSON NOT NULL,
        total DECIMAL(10,2) NOT NULL,
        currencyUnit VARCHAR(3) NOT NULL
    )`)
	if err != nil {
		return err
	}

	return nil
}
