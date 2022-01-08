start:
	docker-compose up -d && \
	docker-compose exec learning-go go run main.go

restart:
	docker-compose down && \
	docker-compose build --no-cache && \
	docker-compose up -d && \
	docker-compose exec learning-go go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

shell:
	docker-compose exec learning-go ash

rebuild:
	docker-compose build --no-cache
