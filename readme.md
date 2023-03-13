# Basic Go Library API With Gin-Gonic

## To Create the go.mod file
`go mod init examples/api`

## To download the framework
`go get github.com/gin-gonic/gin`

## To run the server
`go run .`  or    `go run main.go`

## Add a book
We'll create a json file with the book object in it.
To add the file to the books we'll write this command:

`curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"`

## To 'loan' a book

`curl localhost:8080/checkout?id=2 --request "PATCH" `

## To 'return' a book

`curl localhost:8080/return?id=2 --request "PATCH"`
