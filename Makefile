.PHONY: build run
build:
	go build -o build/VShareServer main.go
run: build
	build/VShareServer
