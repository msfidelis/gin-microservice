.PHONY: build local

local:
	dep ensure
	go build -ldflags="-s -w" -o bin/app main.go

build-docker:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/app main.go

clean:
	rm -rf ./bin
