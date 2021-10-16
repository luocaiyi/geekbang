export tag=v1.0
root:
	export ROOT=github.com/luocaiyi/geekbang/CloudNativeTrainingCamp/homework/module3

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t caiyiluo/httpserver:v1.0 .

push: release
	echo "pushing caiyiluo/httpserver"
	docker push caiyiluo/httpserver:v1.0
