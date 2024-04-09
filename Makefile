start:
	go run cmd/sso/main.go --config=./config/local.yml
migrations:
	go run cmd/migrations/migrations.go