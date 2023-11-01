# Day 20

## Building a Basic Web Application

This is a simple Go web application that demonstrates basic CRUD operations for a simple data entity. The application uses the `net/http` package to handle HTTP requests and responses. It serves as a basic template for implementing CRUD functionality in a web application.

## Getting Started

To run the application, make sure you have Go installed on your system. You can then run the following command:

```bash
go run main.go
```

The server will start at http://localhost:8080.

## Endpoints

- `GET /items` - Retrieve all items
- `POST /items/create` - Create a new item
- `PUT /items/update` - Update an existing item
- `DELETE /items/delete` - Delete an item

## Data Structure

The `Item` struct represents the data entity used in the application. It has the following properties:

```go
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

## Usage

You can use tools like cURL or Postman to send requests to the defined endpoints. Ensure you provide valid JSON data for creating and updating items.

### Example Usage

#### Create an item:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "name": "Sample Item"}' http://localhost:8080/items/create
```

#### Get all items:

```bash
curl http://localhost:8080/items
```

#### Update an item:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "name": "Updated Item"}' http://localhost:8080/items/update
```

#### Delete an item:

```bash
curl -X DELETE http://localhost:8080/items/delete
```
