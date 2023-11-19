v=?

all: install

publish:
	@git tag v$(v)
	@git push origin v$(v)

install:
	@go mod tidy

test:
	@go test -coverprofile=coverage.out -v ./...
	@go tool cover -html=coverage.out

.PHONY: all \
		publish \
		install	\
		test
