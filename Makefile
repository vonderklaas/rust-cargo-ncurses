MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migration-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@, $(MAKECMDGOALS))