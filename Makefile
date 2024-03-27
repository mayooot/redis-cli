

build:
	GOOS=linux GOARCH=amd64 go build -o redis-cli-linux-amd64 redis-cli.go
	GOOS=darwin GOARCH=amd64 go build -o redis-cli-darwin-amd64 redis-cli.go
 	GOOS=windows GOARCH=amd64 go build -o redis-cli-windows-amd64.exe redis-cli.go
