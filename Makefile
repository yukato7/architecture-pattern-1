.PHONY: go

setup:
		@if [ ! "$(shell go version|awk  '{print $3}' | grep go1.11)" ]; then \
				echo "Go version 1.11 ~ is  required."; \
				exit 1; \
		fi

build: setup
		@echo "build"
		@GO111MODULE=on go get -u ./
		@GO111MODULE=on GOOS=linux GOARCH=amd64  CGO_ENABLED=0 go build -o ./bin/api  ./

dbuild:
		@echo "=== Docker build ==="
		@docker-compose build
dstart: 
		@echo  "=== Docker start ==="
		@docker-compose up -d
dlog:
		@docker ps --filter name="gengo-api" --format "{{.Names}}" | peco | xargs docker logs -f