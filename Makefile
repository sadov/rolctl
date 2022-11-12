all: generate

generate: 
	go mod init github.com/sadov/rolctl
	swagger generate cli -f ../rol/src/webapi/swagger/swagger.json -A rol-cmd --skip-validation
	go mod tidy

modules:
	cd models && go mod init github.com/sadov/rolctl/models && go mod tidy
	cd cli &&  go mod init github.com/sadov/rolctl/cli && go mod tidy
	cd client && go mod init github.com/sadov/rolctl/client && go mod tidy
	cd cmd && go mod init github.com/sadov/rolctl/cmd && go mod tidy
	go mod tidy

clean:
	rm -rf cli client cmd go.mod go.sum models
