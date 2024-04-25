DB_URL=postgresql://root:root@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d arm64v8/postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migrateup4:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 4

sqlc:
	sqlc generate

server:
	go run main.go

docker:
	docker-compose up

.PHONY: postgres createdb migrateup migratedown sqlc server docker migrateup4