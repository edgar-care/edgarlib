v=?

all: install

publish:
	@git tag v$(v)
	@git push origin v$(v)

install:
	@go mod tidy

generate:
	@go get github.com/99designs/gqlgen@v0.17.43
	@cd graphql && go run github.com/99designs/gqlgen generate && go run hooks/bson.go && go run client/generate_client.go client/client_generator.go client/parser.go

test:
	@bold=$$(tput bold); \
	normal=$$(tput sgr0); \
	count=$$(go test -coverprofile=coverage.out -v ./appointment ./auth ./dashboard ./follow_treatment | tee /dev/tty | grep -c '=== RUN'); \
	total_coverage=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}'); \
	echo "\n==========================================\n" && \
	echo "$${bold}Number of tests executed: $${count}\n" && \
    echo "Total Coverage: $${total_coverage}$${normal}\n" && \
	echo "=========================================="
	@go tool cover -html=coverage.out

test-server:
	@docker build -t edgar-mongodb-image .
	@docker run --rm --name edgar-mongodb-test -p 27017:27017 -d edgar-mongodb-image
	@DATABASE_URL="mongodb://localhost:27017" go run graphql/test/test_server.go
	@docker stop my-mongodb-test

clean-test-server:
	@mongosh mongodb://localhost:27017 --eval "db.getSiblingDB('web').dropDatabase();"

.PHONY: all \
		publish \
		install	\
		test
