.PHONY: build

CMD_FILE='./cmd/datafob/main.go'

build:
	rm -rf ./build
	mkdir -p ./build
	go build -o ./build/main ${CMD_FILE}
	echo ./build/main

run:
	go run ${CMD_FILE}


