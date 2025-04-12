.PHONY: up down build logs restart shell

up:
	docker compose -f .docker/docker-compose.yml up -d

down:
	docker compose -f .docker/docker-compose.yml down

build:
	docker compose -f .docker/docker-compose.yml build

debugbuild:
	docker compose -f .docker/docker-compose.yml build --no-cache --progress=plain

logs:
	docker compose -f .docker/docker-compose.yml logs -f

reup: down up

shell:
	docker exec -it gb_tgbot-app bash

# Run only app container for debugging
debugshell:
	docker compose -f .docker/docker-compose.yml run --rm -it app sh

dev:
	docker compose -f .docker/docker-compose.yml up --build

# usage: make newmigration name=hello_new_table
newmigration:
	docker compose -f .docker/docker-compose.yml exec app migrate create -ext sql -dir migrations -seq "$(name)"

# runs migrations
migrate:
	docker compose -f .docker/docker-compose.yml exec app migrate -path=migrations -database "mysql://user:pass@tcp(db:3306)/mydb?multiStatements=true" up


# rollbacks migrations
unmigrate:
	docker compose -f .docker/docker-compose.yml exec app migrate -path=migrations -database "mysql://user:pass@tcp(db:3306)/mydb?multiStatements=true" down