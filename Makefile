v=?

all: install

publish:
	@git tag v$(v)
	@git push origin v$(v)

install:
	@go mod tidy

generate:
	@go get github.com/Khan/genqlient/generate@v0.6.0
	@cd graphql && go run github.com/Khan/genqlient && go run github.com/99designs/gqlgen generate

test:
	@go test -coverprofile=coverage.out -v ./...
	@go tool cover -html=coverage.out

.PHONY: all \
		publish \
		install	\
		test
