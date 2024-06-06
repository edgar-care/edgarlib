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
	count=$$(grc go test -coverprofile=coverage.out -v ./appointment ./auth ./dashboard ./follow_treatment ./chat ./black_list ./medicament ./redis ./slot ./treatment | tee /dev/tty | grep -c '=== RUN'); \
	total_coverage=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}'); \
	echo "\n==========================================\n" && \
	echo "$${bold}Number of tests executed: $${count}\n" && \
    echo "Total Coverage: $${total_coverage}$${normal}\n" && \
	echo "=========================================="
	@go tool cover -html=coverage.out

test-server:
	@docker build -t edgar-mongodb-image -f Dockerfile.mongo .
	@docker run --rm --name edgar-mongodb-test -p 27017:27017 -d edgar-mongodb-image

	@docker build -t edgar-redis-image -f Dockerfile.redis .
	@docker run --rm --name edgar-redis-test -p 6379:6379 -p 2222:22 -d edgar-redis-image

	@docker run --rm --name edgar-mailhog -p 1025:1025 -p 8025:8025 -d mailhog/mailhog

	@DATABASE_URL="mongodb://localhost:27017" go run graphql/test/test_server.go
	@docker stop my-mongodb-test
	@docker stop edgar-redis-test

stop-test-dockers:
	@docker stop edgar-mongodb-test
	@docker stop edgar-redis-test
	@docker stop edgar-mailhog

clean-test-server:
	@mongosh mongodb://localhost:27017 --eval "db.getSiblingDB('web').dropDatabase();"

.PHONY: all \
		publish \
		install	\
		test
