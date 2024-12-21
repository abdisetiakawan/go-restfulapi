# Go RESTful API

This project is a simple RESTful API built with Go, intended for learning purposes. It demonstrates basic CRUD operations on a `Category` resource.

## Project Structure

- `internal/service`: Contains the service layer where business logic is implemented.
- `internal/repository`: Contains the repository layer for database interactions.
- `internal/model`: Contains the data models.
- `internal/helper`: Contains helper functions.

## Prerequisites

- Go 1.16 or later
- A running instance of a SQL database (e.g., MySQL, PostgreSQL)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/abdisetiakawan/go-restfulapi.git
    cd go-restfulapi
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your database and update the database configuration in the project.

## Usage

### Running the Server

To start the server, run:
```sh
go run main.go
```

### API Endpoints

- **Create Category**
    - **URL**: `/categories`
    - **Method**: `POST`
    - **Request Body**:
        ```json
        {
            "name": "Category Name"
        }
        ```
    - **Response**:
        ```json
        {
            "id": 1,
            "name": "Category Name"
        }
        ```

- **Update Category**
    - **URL**: `/categories/{id}`
    - **Method**: `PUT`
    - **Request Body**:
        ```json
        {
            "name": "Updated Category Name"
        }
        ```
    - **Response**:
        ```json
        {
            "id": 1,
            "name": "Updated Category Name"
        }
        ```

- **Delete Category**
    - **URL**: `/categories/{id}`
    - **Method**: `DELETE`
    - **Response**: `200`

- **Get Category by ID**
    - **URL**: `/categories/{id}`
    - **Method**: `GET`
    - **Response**:
        ```json
        {
            "id": 1,
            "name": "Category Name"
        }
        ```

- **Get All Categories**
    - **URL**: `/categories`
    - **Method**: `GET`
    - **Response**:
        ```json
        [
            {
                "id": 1,
                "name": "Category Name"
            },
            {
                "id": 2,
                "name": "Another Category Name"
            }
        ]
        ```

## Contributing

Feel free to fork this repository and contribute by submitting a pull request.
