up:
	docker-compose up -d
build:
	docker-compose build --no-cache --force-rm
create-project:
	@make build
	@make up
start:
	docker-compose start
stop:
	docker-compose stop
down:
	docker-compose down --remove-orphans
restart:
	@make down
	@make up
destroy:
	docker-compose down --rmi all --volumes --remove-orphans
app:
	docker-compose exec app bash
db:
	docker-compose exec db bash