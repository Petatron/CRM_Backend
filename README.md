# CRM_Backend
## Project Specifications

### Usage
1. Installation:
    - Clone the repository
    - Run `go get` to install dependencies
    - Run `go run main.go` to start the server
    - Navigate to `localhost:3000` in your browser
    - Use the API endpoints to interact with the database
2. API Endpoints:
    - `/` - Home route
    - `/customers` - Get all customers
    - `/customers/{id}` - Get a single customer
    - `/customers` - Add a customer
    - `/customers/{id}` - Update a customer
    - `/customers/{id}` - Delete a customer

### Data
1. Customer struct with the following fields:
    - ID
    - Name
    - Role
    - Email
    - Phone
    - Contacted (i.e., indication of whether the customer has been contacted)
2. Create a mock "database" to store customer data
    - Customers are stored appropriately in a basic data structure (e.g., slice, map, etc.) that represents a "database."
3. Seed the database with initial customer data
   - The "database" data structure is non-empty. That is, prior to any CRUD operations performed by the user (e.g., adding a customer), the database includes at least three existing (i.e., "hard-coded") customers.
4. Assign unique IDs to customers in the database
   - Each customer in the database has a unique ID. That is, no two customers in the database have the same ID.
5. The database is persistent
    - The database is not reinitialized each time the program is run. That is, if the program is run multiple times, the database retains the same data between runs.

### Server
1. Create a mock RESTful server endpoints that can handle requests to the following routes:
   - Getting a single customer through a `/customers/{id}` path
   - Getting all customers through a `/customers` path
   - Creating a customer through a `/customers` path
   - Updating a customer through a `/customers/{id}` path
   - Deleting a customer through a `/customers/{id}` path
   
   Each RESTful route is associated with the correct HTTP verb.
2. Return JSON in server responses
   - The application leverages the encoding/json package to return JSON values (i.e., not text, etc.) to the user.
3. Return appropriate HTTP status codes
4. Serve static HTML at the home ("/") route
    - The application serves static HTML at the "/" route. That is, when the user navigates to the root of the application, they are served a static HTML page.The home route is a client API endpoint, and includes a brief overview of the API (e.g., available endpoints). 
    - Note: This is the only route that does not return a JSON response.
5. Set up and configure a router
    - The application uses a router (e.g., `gorilla/mux`, `http.ServeMux`, etc.) that supports HTTP method-based routing and variables in URL paths.
6. Set up and configure a logger
7. Create and assign handlers for requests
   - The Handler interface is used to handle HTTP requests sent to defined paths. There are five routes that return a JSON response, and are each is registered to a dedicated handler:
      - `getCustomers()`
      - `getCustomer()`
      - `addCustomer()`
      - `updateCustomer()`
      - `deleteCustomer()`
8. Includes basic error handling for non-existent customers.
   - If a user attempts to access a customer that does not exist, the application returns an appropriate error message (e.g., "Customer not found") and an HTTP status code of 404. null or an empty JSON object literal or an error message. 
9. Set headers to indicate the proper media type
    - The application sets the Content-Type header to indicate the proper media type (e.g., application/json) for each response.
10. Read request data
    - The application leverages the `io/ioutil` package to read I/O (e.g., request) data.
11. Parse JSON data
    - The application leverages the `encoding/json` package to parse JSON data.

### Future Improvements
1. Create an additional endpoint that updates customer values in a batch (i.e., rather than for a single customer).
2. Make Customers data private
3. Upgrade the mock database to a real database (e.g., PostgreSQL).
4. Deploy the API to the web.




