
# Udacity: CRM-Backend

						
## Description


This is a build of the backend (i.e., server-side portion) of a CRM application. 
While the server is on and active, users will be able to make HTTP requests to the server to perform CRUD operations.
Since this project doesn't have a user interface for users to interact with, users will be largely interacting with the project via Postman (or cURL).
							

## Installation Instructions

**Clone the repository**:

```bash
 git clone https://github.com/yourusername/yourrepository.git 
```
**Ensure Go is installed**:
 * Make sure you have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/). 
 
**Download dependencies**:
 - Run the following command to download the necessary dependencies specified in the `go.mod` file.
 ```bash
go mod tidy
 ```
 
## Launch
							   
- The application only requires a simple `go run` command to launch:
```bash
go run main.go
 ```
							
## Usage

- When you send a simple "GET" request to "localhost:3000", you will be able to see all the available paths.

You can also :

- Create a customer using: POST in the "localhost:3000/customers" path.

- Delete a customer using: DELETE in the "localhost:3000/customers/{id}" path.

- Update a customer using: PUT in the "localhost:3000/customers/{id}" path.

- Get a single customer using: GET in the "localhost:3000/customers/{id}" path.

- Get all the customers using: GET in the "localhost:3000/customers" path
