build:
	go build -o ./bin/task ./cmd/main.go
migrateup:
	export $(grep -v '#.*' .env | xargs) && migrate -database ${POSTGRESQL_URL} -path migrations up
migratedown:
	export $(grep -v '#.*' .env | xargs) && migrate -database ${POSTGRESQL_URL} -path migrations down
start:
	docker-compose up --build