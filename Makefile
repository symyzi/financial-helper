.PHONY: createdb dropdb postgres migrateup migratedown dockerstart dockerstop sqlc test

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root financial_helper

dropdb:
	docker exec -it postgres16 dropdb financial_helper

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/financial_helper?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/financial_helper?sslmode=disable" -verbose down

dockerstart:
	docker start postgres16

dockerstop:
	docker stop postgres16

sqlc:
	sqlc generate

test:
	go test -v -cover ./...