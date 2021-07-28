run:
	docker-compose up

update-app:
	docker-compose up -d --no-deps --build app

update-db:
	docker-compose up -d --no-deps --build db

down:
	docker-compose down