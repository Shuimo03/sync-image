build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o sync-image cmd/sync-image/sync-image.go
