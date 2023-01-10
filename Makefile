.PHONY: dependencies

run:

dependencies:
	docker compose -f scripts/docker-compose.yml up -d
