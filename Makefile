GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	cd lambda; GOOS=linux GOARCH=amd64 go build -o hello/main ./hello
	cd lambda; GOOS=linux GOARCH=amd64 go build -o list-tables/main ./list-tables

start: fmt build
	docker network inspect aws-local &>/dev/null || \
		docker network create aws-local
	docker-compose up -d

stop:
	docker-compose down

restart: fmt build
	docker-compose restart
