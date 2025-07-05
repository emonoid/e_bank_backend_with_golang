postgres16:
	docker run --name postgres16 -p 5454:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.9-bullseye

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root islami_bank

dropdb:
	docker exec -it postgres16 dropdb islami_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres16 migrateup migratedown sqlc test server migrateup1 migratedown1
