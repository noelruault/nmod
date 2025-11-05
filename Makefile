# Default port
PORT ?= 9000

# Run the API server
run:
	PORT=$(PORT) go run -C api main.go

# Tidy dependencies for both modules
tidy:
	cd iprange && go mod tidy
	cd api && go mod tidy
