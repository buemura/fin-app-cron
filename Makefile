build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o main cmd/main.go
zip:
	zip main.zip main