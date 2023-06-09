VERSION=
BINARY_NAME=./build/secure-pass

build:
	@if [ -z "${VERSION}" ]; then \
		echo "VERSION is not defined. Please define VERSION using the command 'make VERSION=<version>'"; \
		exit 1; \
	fi
	echo "Building version ${VERSION}"
	./keygen
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-${VERSION}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-${VERSION}-linux-amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-${VERSION}-windows-amd64 main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-${VERSION}-darwin-arm64 main.go
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}-${VERSION}-linux-arm64 main.go
	./clearkey

clear:
	rm -r build

test:
	go clean -testcache && go test -v ./...