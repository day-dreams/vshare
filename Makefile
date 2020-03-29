.PHONY: build run


WORKDIR=$(CURDIR)

build:
	go build -o build/VShareServer main.go
run: build
	build/VShareServer --vfile=$(WORKDIR)/runtime/videos.json
