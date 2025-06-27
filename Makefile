postgres16:
	docker run --name postgres16 -p 5454:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.9-bullseye

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root islami_bank

dropdb:
	docker exec -it postgres16 dropdb islami_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres16 migrateup migratedown
