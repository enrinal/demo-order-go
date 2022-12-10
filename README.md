# Demo Order API Using GO

## Description
This is an example of implementation of Clean Architecture in Go (Golang) projects. But with some case :)

Rule of Clean Architecture by Uncle Bob

* Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
* Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
* Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
* Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

This project has 4 Domain layer :

* Models Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

#### The Diagram
![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

#### The Case
* Buyers can only buy a maximum of 5 times a day
* Buyers can buy goods and buy a maximum of 10 items per product
* Buyers can enter into the Basket before processing the purchase
* The buyer processes the purchase of the order
* Buyer will see pending purchase status
* The seller will process the order and the purchase status is processed
* Send goods and purchase status to be sent
* Buyers can see the status of the purchase whether it is already in process, has been sent or is still pending
* Buyers can change the status to be accepted or automatically changed if it passes 1 day.

This project using Mysql and MongoDb

### API List

#### Order
* GET /api/v1/orders
```json
[
    {
        "orderId": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
        "buyer_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
        "status": "pending",
        "created_at": "2019-08-01T00:00:00Z",
        "updated_at": "2019-08-01T00:00:00Z",
        "items": [
            {
                "id": 1,
                "quantity": 1
            }
        ]
    }
]
```
* GET /api/v1/orders/:id
```json
{
    "order_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "buyer_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "status": "pending",
    "created_at": "2019-08-01T00:00:00Z",
    "updated_at": "2019-08-01T00:00:00Z",
    "items": [
        {
            "id": 1,
            "quantity": 1
        }
    ],
    "total_price": 100000,
    "total_quantity": 1,
    "total_item": 1
}
```
* POST /api/v1/orders
```json
{
    "buyer_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "cart_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a"
}
```
* POST /api/v1/orders/:id/status
```json
{
    "status": "pending"
}
```

#### Product
* GET /api/v1/products
```json
[
    {
        "id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
        "name": "Product 1",
        "price": 10000,
        "created_at": "2019-08-01T00:00:00Z",
        "updated_at": "2019-08-01T00:00:00Z"
[...]
]
```
* GET /api/v1/products/:id
```json
{
    "id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "name": "Product 1",
    "price": 10000,
    "created_at": "2019-08-01T00:00:00Z",
    "updated_at": "2019-08-01T00:00:00Z"
}
```

#### User
* POST /api/v1/users
```json
{
  "email": "enrinal@gmail.com",
  "password": "123456"
}
```
* POST /api/v1/users/login
  - Request
```json
{
  "email": "enrinal@gmail.com",
  "password": "123456"
}
```
  - Response
```json
{
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW5yaW5hbCIsInBob25lIjoiNjI4MTI4MTc5NDg0OSIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTY1NDk2Mzc4MX0.Ek8E8uMmbIqlNhJ6Q5G-9xN8vI8kRMhwXno89CeuCh8"
  }
}
```
* GET /api/v1/users/claims
```json
{
  "message": "success",
  "data": {
    "name": "enrinal",
    "phone": "6281281794849",
    "role": "admin"
  }
}
```

#### Cart
* GET /api/v1/carts/:id
```json
{
    "id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "buyer_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "created_at": "2019-08-01T00:00:00Z",
    "updated_at": "2019-08-01T00:00:00Z",
    "products": [
        {
            "id": 1,
            "quantity": 1
        }
    ]
}
```
* POST /api/v1/carts
```json
{
    "buyer_id": "5e4f3c7c-3f8a-4c9e-8c1a-3b9f3c7c3f8a",
    "products": [
        {
            "id": 1,
            "quantity": 1
        }
    ]
}
```

### How To Run This Project
>  Make Sure you have run the simpleorder.sql in your mysql

```bash
# Clone into YOUR $GOPATH/src
git clone https://github.com/enrinal/demo-order-go.git

# Build Project
go build

# Run Project
./simple-order
```
