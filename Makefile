running-dev:
	@nodemon -x go run main.go
dev:
	@go run main.go
	
build:
	@go build -o main