.PHONY: ent


install:
	go install github.com/google/wire/cmd/wire@latest

serve:
	reflex -r '\.go$$' -R '(^vendor/.*)' -s -- sh -c 'clear && go run ./cmd/serve'

build:
	go build ./cmd/serve -o ./bin

env:
	cp ./sample.env ./.env

wire:
	wire ./di

# Volume actions
clean_volumes:
	sudo rm -rf ./tmp/redis ./tmp/postgres ./tmp/localstack

create_db_dump:
	docker exec arrndme_postgres pg_dump -U arrndme -F c arrndme > arrndme.dump

restore_db:
	docker cp  mydb.dump arrndme_postgres:/ 
	docker exec arrndme_postgres pg_restore -U arrndme  -d arrndme  arrndme.dump

ent:
	go generate ./ent

ent_generate:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

generate_migrations:
	atlas migrate diff migration_name \
  		--dir "file://ent/migrate/migrations" \
  		--to "ent://ent/schema" \
		--dev-url "postgresql://$(DB_USER):$(DB_PWD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?search_path=public&sslmode=disable"

migrate:
	atlas migrate apply \
  		--dir "file://ent/migrate/migrations" \
		--url "postgresql://$(DB_USER):$(DB_PWD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?search_path=public&sslmode=disable"