package models

type CommonAPIResponse struct {
	IsError         bool
	ResponseMessage string
	ResponseData    interface{}
}

type Order struct {
	ID           string      `json:"id" db:"id"`
	Status       string      `json:"status" db:"status"`
	Items        []OrderItem `json:"items" db:"items"`
	Total        float64     `json:"total" db:"total"`
	CurrencyUnit string      `json:"currencyUnit" db:"currencyUnit"`
}

type OrderItem struct {
	ID          string  `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int     `json:"quantity" db:"quantity"`
}

type CreateOrderRequest struct {
	Items        []OrderItem `json:"items" binding:"required"`
	Total        float64     `json:"total" binding:"required"`
	CurrencyUnit string      `json:"currencyUnit" binding:"required"`
	Status       string      `json:"status"`
}

type UpdateOrderRequest struct {
	Status string `json:"status" binding:"required"`
}
