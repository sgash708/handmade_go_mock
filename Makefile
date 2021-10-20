ls:
	docker ps -a
stats:
	docker stats
logs_server:
	docker logs go_server
logs_db:
	docker logs mock_mysql
images:
	docker images
rmi:
	docker rmi ${ARG}
build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down
exec:
	docker-compose exec go_server bash