postgres:
	docker run --name DB_ECOMMERCE -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it DB_ECOMMERCE createdb --username=root --owner=root ECOMMERCE

dropdb:
	docker exec -it DB_ECOMMERCE dropdb ECOMMERCE

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:0945639220Beo@localhost:5432/ECOMMERCE?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:0945639220Beo@localhost:5432/ECOMMERCE?sslmode=disable" -verbose down

migrateup1:
	migrate -path internal/db/migration -database "postgresql://root:0945639220Beo@localhost:5432/ECOMMERCE?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path internal/db/migration -database "postgresql://root:0945639220Beo@localhost:5432/ECOMMERCE?sslmode=disable" -verbose down 1

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
	mockgen -package mock -destination internal/db/mock/mock.go github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc Querier

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server wire mock