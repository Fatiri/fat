sqlc:
	sqlc generate
swagger:
	swag init -g api/api.go
mock:
	mockery --all --keeptree	
generate_migration:
	migrate create -ext sql -dir ./repository/migrations -seq $(table)
dbup:
	migrate -database "postgres://$(username):$(password)@localhost:5432/trading_fat?sslmode=disable" -path ./repository/migrations up
dbdown:
	migrate -database "postgres://$(username):$(password)@localhost:5432/trading_fat?sslmode=disable" -path ./repository/migrations down
dbrollback:
	migrate -path ./repository/migrations -database "postgres://$(username):$(password)@localhost:5432/trading_fat?sslmode=disable" force $(version)
run_dev_cli:
	go run . --environment=dev --service=CLI
run_stg_cli:
	go run . --environment=staging --service=CLI
run_prod_cli:
	go run . --environment=production --service=CLI
run_dev_api:
	go run . --environment=dev --service=API
run_stg_api:
	go run . --environment=staging --service=API
run_prod_api:
	go run . --environment=production --service=API
run_dev_gui:
	go run . --environment=dev --service=GUI
run_stg_gui:
	go run . --environment=staging --service=GUI
run_prod_gui:
	go run . --environment=production --service=GUI
test:
	go test -coverprofile=cover.out ./... && go tool cover -func=cover.out && go tool cover -html=cover.out