dev-serve:
	@nodemon -x go run main.go
dev:
	@go run main.go
build:
	@go build -o villa-go .
tidy:
	@go mod tidy