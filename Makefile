run:
	go run main.go

migrateup:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/product" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/product" -verbose down

.PHONY: migrateup migratedown