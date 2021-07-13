run:
	go build -o bookings cmd/web/*.go && ./bookings
test:
	go test -v ./...