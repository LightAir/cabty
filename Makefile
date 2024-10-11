.DEFAULT_GOAL: lint

BIN_PATH_NAME=./bin/cabty
GOLANG_LINT_CMD=golangci-lint run

.PHONY: *

build:
	GOARCH=amd64 GOOS=linux go build -o ${BIN_PATH_NAME} main.go

run: build
	${BIN_PATH_NAME}

clean:
	go clean
	rm ./bin/*

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.61.0 ${GOLANG_LINT_CMD}

lint-bin:
	${GOLANG_LINT_CMD}