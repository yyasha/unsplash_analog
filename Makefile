gen_docs:
	swag init http_api/*

migrate_create:
	migrate create -ext sql -dir postgres/migrations -seq ${NAME}

docker_build:
	docker build -t registry.computernetthings.ru/unsplash/backend:latest . && docker image prune -f

docker_push:
	docker push registry.computernetthings.ru/unsplash/backend:latest