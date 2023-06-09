# Orders API

This is a simple RESTful API for managing orders. It allows you to create, retrieve, and update orders.

## Technologies Used
This API is built using the following technologies:

* Golang
* Gin
* SQLx
* MySQL
* UUID

## Installation
```shell
docker run --name mysql -d \
    -p 3307:3306 \
    -e MYSQL_ROOT_PASSWORD=root \
    --restart unless-stopped \
    mysql:8
export SQL_URL="root:root@tcp(localhost:3307)/mysql"
go run main.go
```

## API Endpoints

**Create Order**
`POST /orders`

Request Body

* items (required) - An array of order items.
* total (required) - The total cost of the order.
* currencyUnit (required) - The currency unit for the total.
* status - The status of the order. Defaults to "pending".

```shell
curl --location 'http://localhost:8080/api/v1/orders' \
--header 'Content-Type: application/json' \
--data '{
    "items": [
        {
            "description": "Product A",
            "price": 10.0,
            "quantity": 2
        },
        {
            "description": "Product B",
            "price": 5.0,
            "quantity": 3
        }
    ],
    "total": 35.0,
    "currencyUnit": "USD",
    "status": "processing"
}'

```

Response Body

* id - The ID of the newly created order.

```shell
{
    "id": "1234567890"
}
```



**Get Order**
`GET /orders/:id`

Request Body

* id (required) - The ID of the order to retrieve.

```shell
curl --location 'http://localhost:8080/api/v1/orders'
```

Response Body

* id - The ID of the order.
* status - The status of the order.
* items - An array of order items.
* total - The total cost of the order.
* currencyUnit - The currency unit for the total.

```shell
{
    "id": "1234567890",
    "status": "processing",
    "items": [
        {
            "id": "0987654321",
            "description": "Product A",
            "price": 10.0,
            "quantity": 2
        },
        {
            "id": "1122334455",
            "description": "Product B",
            "price": 5.0,
            "quantity": 3
        }
    ],
    "total": 35.0,
    "currencyUnit": "USD"
}
```


**Get Orders with sort**
`GET /orders`

Query with
```shell
curl --location 'http://localhost:8080/api/v1/orders?status=processing'
```
* status - Filter orders by status.
* sortBy - Sort orders by field.
* sortOrder - Sort order direction. Defaults to "asc".


Response Body

* id - The ID of the order.
* status - The status of the order.
* items - An array of order items.
* total - The total cost of the order.
* currencyUnit - The currency unit for the total.

```shell
{
    "id": "1234567890",
    "status": "processing",
    "items": [
        {
            "id": "0987654321",
            "description": "Product A",
            "price": 10.0,
            "quantity": 2
        },
        {
            "id": "1122334455",
            "description": "Product B",
            "price": 5.0,
            "quantity": 3
        }
    ],
    "total": 35.0,
    "currencyUnit": "USD"
}
```


**Update Order**
`PUT /orders/:id`

```shell
curl --location --request PUT 'http://localhost:8080/api/v1/orders/:id' \
--header 'Content-Type: application/json' \
--data '{
    "status": "done"
}'
```
Request Body

* status - Filter orders by status.


Response Body

* id - The ID of the order.

```shell
{
    "id": "1234567890"
}
```