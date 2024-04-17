start:
	go run cmd/sso/main.go --config=./config/local.yml
migrations:
	go run cmd/migrations/migrations.go --config=./config/local.yml
up:
	cd ./docker && docker compose up -d