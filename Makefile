all: rolctl

go.mod go.sum:
	go mod init github.com/sadov/rolctl
	swagger generate cli -f ../rol/src/webapi/swagger/swagger.json -A rol-cmd --skip-validation
	go mod tidy

rolctl: go.mod go.sum
	go build -o rolctl cmd/cli/main.go

clean:
	rm -rf rolctl cli client cmd go.mod go.sum models
