APPNAME=gophr.linux

build: mod
	go build -o ${APPNAME}

mod:
	go mod download

fmt:
	go fmt ./...

docker:
	sudo docker build -t gophr .

run-step: run clean

run: docker
	docker-compose up -d --no-deps

start:
	docker-compose start

stop:
	docker-compose stop

down:
	docker-compose down

clean:
	sudo docker image prune -af

.PHONY: run build mod fmt docker clean down

