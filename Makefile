build: 
	@go build -o bin/exchange

run: build
	@./bin/exchange
