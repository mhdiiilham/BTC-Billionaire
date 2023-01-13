.PHONY: dependencies migrate-create

run:

test:
	go test -race -v ./...

dependencies:
	docker compose -f scripts/docker-compose.yml up -d

migrate-create:
	@read -p  "Migration name (eg:create_users, alter_entities, ...): " NAME; \
	migrate create -ext sql -dir migrations -seq $$NAME

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down -all

mock-prepare:
	go install github.com/golang/mock/mockgen@v1.6.0

mocks:
	mockgen -source=service/interface.go -destination=service/mock/interfaces.go -package=mock

