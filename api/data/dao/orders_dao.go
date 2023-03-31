package dao

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"order-management-service/api/models"
	"os"
)

var OrdersDao OrdersDaoInterface = &ordersDao{}

type OrdersDaoInterface interface {
	Connect() *sqlx.DB
	CreateOrder(orderID string, createOrderRequest models.CreateOrderRequest) error
	GetAllOrders(status string, sortBy string, sortOrder string) ([]models.Order, error)
	GetOrdersByOrderID(orderId string) ([]models.OrderItem, error)
	UpdateOrder(orderId string, order models.Order, req models.UpdateOrderRequest) error
	GetOrder(orderId string, order models.Order) (models.Order, error)
}

type ordersDao struct {
	db *sqlx.DB
}

func (o *ordersDao) CreateOrder(orderID string, createOrderRequest models.CreateOrderRequest) error {
	tx := o.db.MustBegin()
	defer func() {
		if r := recover(); r != nil {
			rollBackError := tx.Rollback()
			if rollBackError != nil {
				return
			}
		}
	}()

	_, err := tx.Exec(`INSERT INTO orders (id, status, total, currencyUnit) VALUES (?, ?, ?, ?)`, orderID, createOrderRequest.Status, createOrderRequest.Total, createOrderRequest.CurrencyUnit)
	if err != nil {
		rollBackError := tx.Rollback()
		if rollBackError != nil {
			return rollBackError
		}
		return err
	}

	for _, item := range createOrderRequest.Items {
		itemID := uuid.NewV4().String()
		_, err := tx.Exec(`INSERT INTO order_items (id, order_id, description, price, quantity) VALUES (?, ?, ?, ?, ?)`, itemID, orderID, item.Description, item.Price, item.Quantity)
		if err != nil {
			rollBackError := tx.Rollback()
			if rollBackError != nil {
				return rollBackError
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (o *ordersDao) Connect() *sqlx.DB {
	url := os.Getenv("SQL_URL")
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil
	}
	o.db = db

	return db
}

func (o *ordersDao) GetOrdersByOrderID(orderId string) ([]models.OrderItem, error) {

	var items []models.OrderItem
	err := o.db.Select(&items, "SELECT id, description, price, quantity FROM order_items WHERE order_id=?", orderId)
	if err != nil {
		logrus.Errorf("%v", err)
		return nil, err
	}

	return items, nil
}

func (o *ordersDao) GetAllOrders(status string, sortBy string, sortOrder string) ([]models.Order, error) {
	query := "SELECT id, status, total, currencyUnit FROM orders"
	args := []interface{}{}

	if status != "" {
		query += " WHERE status=?"
		args = append(args, status)
	}

	if sortBy != "" {
		query += " ORDER BY " + sortBy + " " + sortOrder
	}

	// Query the database for the orders
	var orders []models.Order
	err := o.db.Select(&orders, query, args...)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *ordersDao) UpdateOrder(orderId string, order models.Order, req models.UpdateOrderRequest) error {
	err := o.db.Get(&order, "SELECT id FROM orders WHERE id = ?", orderId)
	if err != nil {
		return err
	}

	// Update the order status
	_, err = o.db.Exec("UPDATE orders SET status = ? WHERE id = ?", req.Status, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (o *ordersDao) GetOrder(orderId string, order models.Order) (models.Order, error) {
	err := o.db.Get(&order, "SELECT id, status, total, currencyUnit FROM orders WHERE id=?", orderId)
	if err != nil {
		return models.Order{}, err
	}

	// Query the database for the order items
	var items []models.OrderItem
	err = o.db.Select(&items, "SELECT id, description, price, quantity FROM order_items WHERE order_id=?", orderId)
	if err != nil {
		return models.Order{}, err
	}

	// Set the order items on the order struct
	order.Items = items
	return order, nil
}
