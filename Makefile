docker/local/db/up:
	@echo "============= starting db locally ============="
	go mod tidy
	docker-compose -f resources/docker/database/docker-compose.yaml up postgres_northwind