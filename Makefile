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

# migrateup1:
# 	migrate -path db/migration -database "postgresql://root:0945639220Beo@localhost:5432/Ecommerce?sslmode=disable" -verbose up

# migratedown1:
# 	migrate -path db/migration -database "postgresql://root:0945639220Beo@localhost:5432/Ecommerce?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd/server/main.go

wire:
	cd internal/wire && wire

# mock:
# 	mockgen -package mockstore -destination db/mock/store.go github.com/NghiaLeopard/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server wire