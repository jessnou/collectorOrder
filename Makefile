docker:
	docker run --name collectorOrder -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	docker exec -it collectorOrder createdb --username=root --owner=root collectorOrder

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/collectorOrder?sslmode=disable" -verbose up

start:
	go run cmd/main.go 10,11,14,15

setup1: docker
setup2: createdb migrateup start