.PHONY: build run


WORKDIR=$(CURDIR)

build:
	go build -o build/VShareServer main.go
run: build
	addr=0.0.0.0:8080 config=./build/config.json build/VShareServer
daemon: build
	(pkill VShareServer || echo "skip kill" )&& nohup build/VShareServer \
	--vfile=$(WORKDIR)/runtime/config.json \
	--vindex=$(WORKDIR)/runtime/index.html \
	&> ./build/vshare.log &
buildLinux:
	GOOS=linux GOARCH=amd64 go build -o build/VShareServer main.go
docker-compose: buildLinux
	docker-compose up --build --force-recreate
docker: buildLinux
	docker run \
	-p 8080:8080 \
	-v ~/Desktop:/data \
	--rm -it `docker build -q .`
