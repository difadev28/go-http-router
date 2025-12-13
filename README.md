# Go HTTP Router

A learning project to understand HTTP Router usage in Go using the `httprouter` library.

## Description

This project demonstrates basic HTTP routing implementation in Go using the `github.com/julienschmidt/httprouter` library. The project includes:

- Basic routing with GET requests
- URL parameters (path parameters)
- Pattern matching with multiple parameters
- Catch-all parameters for wildcard routes
- File serving with embedded filesystem

## Dependencies

- `github.com/julienschmidt/httprouter v1.3.0` - High performance HTTP router for Go
- `github.com/stretchr/testify v1.9.0` - Testing library (development only)

## Installation

```bash
git clone https://github.com/difadev28/go-http-router.git
cd go-http-router
go mod tidy
```

## Running the Application

```bash
go run main.go
```

The server will run on `http://localhost:3000`

## API Endpoints

### GET /
Displays a welcome message.

**Example Response:**
```
Hello Http Router
```

### GET /products/:id
Displays product information by ID.

**Example Request:**
```
GET /products/1
```

**Example Response:**
```
Product 1
```

### GET /products/:id/items/:itemsId
Displays item information within a product.

**Example Request:**
```
GET /products/1/items/2
```

**Example Response:**
```
Product 1 Items Id : 2
```

### GET /images/*image
Wildcard route for serving images.

**Example Request:**
```
GET /images/small/profile.png
```

**Example Response:**
```
Image :/small/profile.png
```

### GET /files/*filepath
Serves static files from the resources folder.

**Example Request:**
```
GET /files/hello.txt
```

**Example Response:**
```
Hello HtppRouter
```

## Testing

Run all tests:

```bash
go test -v
```

Run specific tests:

```bash
go test -v -run TestRouter
go test -v -run TestRouterParams
go test -v -run TestRouterPattern
go test -v -run TestServeFile
```

## Project Structure

```
go-http-router/
├── main.go                 # Main file to run the server
├── router_test.go          # Tests for basic routing
├── router_params_test.go   # Tests for parameter routing
├── router_pattern_test.go  # Tests for pattern routing
├── serve_file_test.go      # Tests for file serving
├── resources/              # Folder for static files
│   ├── hello.txt          # Example file
│   └── goodbye.txt        # Example file
├── go.mod                  # Go module file
└── README.md              # Project documentation
```

## Code Examples

### Basic Routing

```go
router := httprouter.New()
router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    fmt.Fprint(writer, "Hello Http Router")
})
```

### Routing with Parameters

```go
router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    id := params.ByName("id")
    fmt.Fprintf(writer, "Product %s", id)
})
```

### File Serving

```go
//go:embed resources
var resources embed.FS

directory, _ := fs.Sub(resources, "resources")
router.ServeFiles("/files/*filepath", http.FS(directory))
```

## Contributing

This is a learning project. Contributions are welcome to add features or improve documentation.

## License

MIT License