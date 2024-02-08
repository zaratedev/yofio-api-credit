.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/credit credit/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

run: clean
	cp .env.local .env
	go mod download
	go run credit/main.go