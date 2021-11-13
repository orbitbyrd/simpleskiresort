build:
	docker compose build

run: build
	docker-compose down --volumes
	docker compose up -d --remove-orphans