.PHONY: build run dev

build:
	go build -o ./cmd/main ./cmd

run:
	go run ./cmd/main.go

running:
	CompileDaemon -build="go build -o ./cmd/main ./cmd" -command=./cmd/main