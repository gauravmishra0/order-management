package data

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"order-management-service/api/data/dao"
	"time"
)

func Initialize() error {
	url := "root:root@tcp(localhost:3307)/mysql"
	logrus.Info(url)
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
	r, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    status VARCHAR(20) NOT NULL,
    items JSON NOT NULL DEFAULT (JSON_OBJECT()),
    total DECIMAL(10,2) NOT NULL,
    currencyUnit VARCHAR(3) NOT NULL
)`)
	db.Exec(`CREATE TABLE IF NOT EXISTS order_items (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  order_id VARCHAR(36) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  quantity INT NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
) ENGINE=InnoDB;
`)
	db.Exec(`ALTER TABLE orders CHANGE COLUMN items JSON json NOT NULL DEFAULT (JSON_OBJECT())`)
	fmt.Print(r)
	if err != nil {
		return err
	}

	return nil
}
