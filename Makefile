install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

start:
	docker-compose up -d
	sleep 5
	make db.migration.up
	go run cmd/main.go

stop:
	docker-compose down

db.migration.status:
	cd migrations && goose postgres "user=root password=root dbname=postgres sslmode=disable" status

db.migration.create:
	cd migrations && goose create ${migration_name} sql

db.migration.up:
	cd migrations && goose postgres "user=root password=root dbname=postgres sslmode=disable" up

db.migration.down:
	cd migrations && goose postgres "user=root password=root dbname=postgres sslmode=disable" down