SHELL:=/bin/bash
FILES:=$(shell find . | grep -v 'github\|golang' | grep '\.go')

test: lint
	go test vehicle
	go test slot
	go test parking
	go test commands

lint: ${FILES}
	@for file in ${FILES}; do ./bin/golint $${file}; done

build: test
	@go build -o bin_parking_lot ./src/main.go
	@chmod 755 bin_parking_lot

run: build
	./bin_parking_lot ${FILE}

