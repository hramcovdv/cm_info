BINARY_NAME=cm_info

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

test:
	go test ./...

clean:
	go clean