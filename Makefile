sqlc:
	sqlc generate
swagger:
	swag init -g api/api.go
generate_migration:
	migrate create -ext sql -dir ./repository/migrations -seq $(table)
dbup:
	migrate -database "postgres://postgres:ilhamfatiri@localhost:5432/trading_fat?sslmode=disable" -path ./repository/migrations up
dbdown:
	migrate -database "postgres://postgres:ilhamfatiri@localhost:5432/trading_fat?sslmode=disable" -path ./repository/migrations down
dbrollback:
	migrate -path ./repository/migrations -database "postgres://postgres:ilhamfatiri@localhost:5432/trading_fat?sslmode=disable" force $(version)
server_dev:
	go run . --environment=dev
server_stg:
	go run . --environment=staging
server_prod:
	go run . --environment=production