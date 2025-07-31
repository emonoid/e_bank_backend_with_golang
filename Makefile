postgres16:
	docker run --name postgres16bull -p 5454:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.9-bullseye

createdb:
	docker exec -it postgres16bull createdb --username=root --owner=root islami_bank

dropdb:
	docker exec -it postgres16bull dropdb islami_bank

newmigrate:
	migrate create -ext sql -dir db/migration your_migration_name

migrateup:
	migrate -path db/migration -database "postgresql://root:femon@localhost:2000/islami_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:femon@localhost:2000/islami_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:femon@localhost:2000/islami_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:femon@localhost:2000/islami_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres16 migrateup migratedown sqlc test server migrateup1 migratedown1
