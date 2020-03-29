.PHONY: build run


WORKDIR=$(CURDIR)

build:
	go build -o build/VShareServer main.go
run: build
	build/VShareServer --vfile=$(WORKDIR)/runtime/videos.json
daemon: build
	(pkill VShareServer || echo "skip kill" )&& nohup build/VShareServer --vfile=$(WORKDIR)/runtime/config.json &> ./build/vshare.log &
