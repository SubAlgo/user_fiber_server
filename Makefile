buildDB:
	sh script/createDB.sh

migrate:
	go run migrate/main.go

dev:
	go run serve/main.go
