test:
	go test -v ./...

build-go:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o stripe-app .
