# Go Movies CRUD API

This project is a CRUD API built using Golang. The API is designed to manage a collection of movies without the use of a database. Instead, it utilizes Go structs and slices to perform CRUD operations on the movie data directly within the server.

## Overview

The Movies CRUD API provides endpoints to perform CRUD operations on movies. It is served on `localhost:8080` using the Gorilla Mux router, an external library for routing in Go.

## Routes

The API defines the following routes:

1. **Get All Route**: `/movies` - Retrieve a list of all movies in the collection.
2. **Get By ID Route**: `/movies/{id}` - Retrieve details of a specific movie by its ID.
3. **Create Route**: `/movies` - Create a new movie entry.
4. **Update Route**: `/movies/{id}` - Update details of a specific movie by its ID.
5. **Delete Route**: `/movies/{id}` - Delete a specific movie by its ID.

## Methods

Each route corresponds to a specific function:

- **Get All Route**: `getMovies`
- **Get By ID Route**: `getMovie`
- **Create Route**: `createMovie`
- **Update Route**: `updateMovie`
- **Delete Route**: `deleteMovie`

## Endpoints

There are two types of endpoints:

1. `/movies`: Used for performing operations that do not require a specific movie ID.
2. `/movies/{id}`: Used for operations that require a specific movie ID.

## HTTP Methods

The API uses different HTTP methods for different operations:

- **GET**: Retrieve data with `getMovies` and `getMovie`.
- **POST**: Create new data with `createMovie`.
- **PUT**: Update existing data with `updateMovie`.
- **DELETE**: Remove data with `deleteMovie`.

## Testing

Postman is recommended for testing all the endpoints using the respective HTTP methods.

## Getting Started

To run the API locally:

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Install dependencies by running the command `go get github.com/gorilla/mux`.
4. Build and run the project using the command `go run main.go`.
5. Use Postman to test the endpoints as described above.

## Dependencies

This project uses the Gorilla Mux router for routing HTTP requests. Ensure you have it installed before running the project.

## Contributing

Contributions to this project are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.
