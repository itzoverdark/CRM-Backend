# CRM-Backend

							DESCRIPTION

This is a build of the backend (i.e., server-side portion) of a CRM application. 
While the server is on and active, users will be able to make HTTP requests to the server to perform CRUD operations.
Since this project doesn't have a user interface for users to interact with, users will be largely interacting with the project via Postman (or cURL).
							
							INSTALLATION
							
- For the gorilla/mux package to work, you have to install it first, to do so, you have to type this into the terminal:
 "go mod init" to create a go.mod file. Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module.
        'go mod init example.com/m/v2' to initialize a v2 module.
	
  Then you can install the package by simply typing "go get github.com/gorilla/mux".


							   LAUNCH
							   
- The application only requires a simple "go run" command to launch.
							
							   USAGE

- When you send a simple "GET" request to "localhost:3000", you will be able to see all the available paths.

You can also :

- Create a customer using: POST in the "localhost:3000/customers" path.

- Delete a customer using: DELETE in the "localhost:3000/customers/{id}" path.

- Update a customer using: PUT in the "localhost:3000/customers/{id}" path.

- Get a single customer using: GET in the "localhost:3000/customers/{id}" path.

- Get all the customers using: GET in the "localhost:3000/customers" path
