DATABASE_URL = $(MYSQL_ADDRS)
MIGRATION_DIR = $(shell pwd)/migrations
GOQLGEN = github.com/99designs/gqlgen

.PHONY: migrate
migrate:
	migrate -source "file:$(MIGRATION_DIR)" -database "mysql://$(DATABASE_URL)" $(move) $(step)

gqlgen:
	go run -mod=mod $(GOQLGEN) generate 
