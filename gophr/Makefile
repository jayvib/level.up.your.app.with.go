APPNAME=gophr.linux

build: mod
	go build -race -o ${APPNAME}

mod:
	go mod download

fmt:
	go fmt ./...

docker:
	sudo docker build -t gophr .

run: docker
	docker-compose up -d

.PHONY: run build mod fmt docker

