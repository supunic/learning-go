start:
	docker-compose up -d && \
	docker-compose exec learning-go ash

restart:
	docker-compose down && \
	docker-compose build --no-cache && \
	docker-compose up -d && \
	docker-compose exec learning-go ash

up:
	docker-compose up -d

down:
	docker-compose down

shell:
	docker-compose exec learning-go ash

rebuild:
	docker-compose build --no-cache
