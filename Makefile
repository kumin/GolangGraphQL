DATABASE_URL = $(MYSQL_ADDRS)
MIGRATION_DIR = $(shell pwd)/migrations

.PHONY: migrate
migrate:
	migrate -source "file:$(MIGRATION_DIR)" -database "mysql://$(DATABASE_URL)" $(move) $(step)
