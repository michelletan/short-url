# Load env vars from root .env.dev
ifneq (,$(wildcard ../.env.dev))
	include ../.env.dev
	export
endif

TOOLS_BIN := $(CURDIR)/bin

DATABASE_URL := ${DATABASE_URL}

install-tools:
	GOBIN=$(TOOLS_BIN) go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.0

migrate-up: install-tools
	$(TOOLS_BIN)/migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down: install-tools
	$(TOOLS_BIN)/migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-create:
	$(TOOLS_BIN)/migrate create -ext sql -dir migrations -seq $(name)

db-up:
	docker compose up db -d

db-down:
	docker compose down db

run:
	env $(cat .env | xargs) go run ./cmd/server/main.go