

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://mohirbek:bismillah@localhost:5432/tasks?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://mohirbek:bismillah@localhost:5432/tasks?sslmode=disable' down
