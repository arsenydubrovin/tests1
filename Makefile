run:
	go run .

deps:
	go mod tidy

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

clean:
	rm coverage.html coverage.out