APP=handsongolang
APP_VERSION:=0.1
#APP_COMMIT:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"
HTTP_SERVE_COMMAND="http-serve"
CONFIG_FILE="./.env"

deps:
	go mod download
compile:
	mkdir -p out/
	go build -ldflags "-X main.version=$(APP_VERSION)" -o $(APP_EXECUTABLE) cmd/*.go

build: deps compile

httpserve: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE) $(HTTP_SERVE_COMMAND)