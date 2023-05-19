DATABASE_URL = $(MYSQL_ADDRS)
PROJECT_DIR = $(shell pwd)
APPS_DIR = $(PROJECT_DIR)/apps
GRAPHQL_SCHEMA = $(PROJECT_DIR)/graph/*.graphqls

CMDS = $(shell find $(PROJECT_DIR)/cmd -mindepth 1 -maxdepth 1 -type d)
STATIC_TARGETS = $(addprefix static-,$(CMDS))
BIN_DIR = $(PROJECT_DIR)/bin

MIGRATION_DIR = $(PROJECT_DIR)/migrations
GRAPHQL_LINTER = graphql-schema-linter

GOQLGEN = github.com/99designs/gqlgen
WIREGEN = github.com/google/wire/cmd/wire

.PHONY: migrate
migrate:
	migrate -source "file:$(MIGRATION_DIR)" -database "mysql://$(DATABASE_URL)" $(move) $(step)

gqlgen:
	go run -mod=mod $(GOQLGEN) generate 

graphql-lint:
	$(GRAPHQL_LINTER) -f compact $(GRAPHQL_SCHEMA)

di:
	go run -mod=mod $(WIREGEN) gen $(APPS_DIR)/...

static: $(STATIC_TARGETS)
$(STATIC_TARGETS):
	$(eval FULLPATH="$(subst static-,,$@)")
	$(eval CMD=$(shell basename "${FULLPATH}"))
	env GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(CMD) $(FULLPATH)
