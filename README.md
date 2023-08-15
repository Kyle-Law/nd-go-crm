# Customer Relationship Management (CRM) API

This CRM API is a simple back-end service built in Go that allows users to manage a collection of customer records. It provides functionality for creating, reading, updating, and deleting customer records.

## Description

The API enables the following functionalities:

- **Getting a list of all customers**
- **Getting data for a single customer**
- **Adding a customer**
- **Updating a customer's information**
- **Batch updating customers' information**
- **Removing a customer**
- **API Overview (accessible at the root endpoint)**


## Installation

To install the project, you'll need to have Go installed on your machine. Follow these steps:

1. **Clone the repository:**

 ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
 ```

2. **Build the application:**

 ```bash
   go build -o crm-api
 ```

## Launch

You can launch the API by running the compiled binary:

```bash
./crm-api
```

By default, the server will start on port 8080. You can access the API at `http://localhost:8080`.

## Usage

Here's how you can interact with the API using common HTTP methods:

### Get All Customers

- **Method:** GET
- **Endpoint:** `/customers`
- **Example:**

```bash
curl http://localhost:8080/customers
```

### Get a Single Customer

- **Method:** GET
- **Endpoint:** `/customers/{id}`
- **Example:**

```bash
curl http://localhost:8080/customers/1
```
### Add a Customer

- **Method:** POST
- **Endpoint:** `/customers`
- **Body:**

```json
  {
      "name": "Example Name",
      "role": "Example Role",
      "email": "Example Email",
      "phone": 5550199,
      "contacted": true
  }
```
- **Example:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Example Name", "role": "Example Role", "email": "Example Email", "phone": 5550199, "contacted": true}' http://localhost:8080/customers
```
### Update a Customer

- **Method:** DELETE
- **Endpoint:** `/customers/{id}`
- **Example:**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Updated Name", "role": "Updated Role", "email": "Updated Email", "phone": 5550199, "contacted": true}' http://localhost:8080/customers/1
```

### Delete a Customer

- **Method:** DELETE
- **Endpoint:** `/customers/{id}`
- **Example:**

```bash
curl -X DELETE http://localhost:8080/customers/1
```

### Batch Update Customers

- **Method:** PUT
- **Endpoint:** `/customers/batchUpdate`
- **Body:** An array of customer objects to be updated.

```json
[
  {
      "id": "1",
      "name": "Updated Name",
      "role": "Updated Role",
      "email": "Updated Email",
      "phone": 5550199,
      "contacted": true
  },
  {
      "id": "2",
      "name": "Another Updated Name",
      "role": "Another Updated Role",
      "email": "AnotherUpdatedEmail@example.com",
      "phone": 5550299,
      "contacted": false
  }
]
```
- **Example:** 
```bash
curl -X PUT -H "Content-Type: application/json" -d @batch_update.json http://localhost:8080/customers/batchUpdate
```

## Testing

You can run the unit tests for the project by executing:

```bash
go test
```

## Contributions

Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
