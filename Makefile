.PHONY: up down build logs restart shell

up:
	docker compose -f .docker/docker-compose.yml up

down:
	docker compose -f .docker/docker-compose.yml down

build:
	docker compose -f .docker/docker-compose.yml build

logs:
	docker compose -f .docker/docker-compose.yml logs -f

reup: down up

shell:
	docker exec -it gb_tgbot-app bash

# Run only app container for debugging
debugshell:
	docker compose -f .docker/docker-compose.yml run --rm -it app bash