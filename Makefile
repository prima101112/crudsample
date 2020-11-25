dafault: build-linux

run: build-linux
	 docker-compose up -d

run-bin: 
	go build -o ./bin/crudsample ./cmd/server/main.go
	./bin/crudsample

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/crudsample ./cmd/server/main.go
	docker-compose build --no-cache
