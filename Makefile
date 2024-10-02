DB_URL=postgresql://root:secret@localhost:5432/ECOMMERCE?sslmode=disable
PATH_DB=db/migration

postgres:
	docker run --name DB_ECOMMERCE -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it DB_ECOMMERCE createdb --username=root --owner=root ECOMMERCE

dropdb:
	docker exec -it DB_ECOMMERCE dropdb ECOMMERCE

migrateup:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd/server/main.go

initDB:
	go run ./cmd/server/initDB.go

wire:
	cd internal/wire && wire

mock:
	mockgen -package mock -destination db/mock/mock.go github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc Querier

swagger:
	swag init --parseDependency --parseInternal -g cmd/server/main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server initDB wire mock