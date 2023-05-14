runapp:
	docker start zhiznmart
	go run cmd/main.go

pullmysql:
	docker pull mysql:latest

runmysql:
	docker run --name zhiznmart -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:latest

createdb:
	docker exec -it zhiznmart mysql -u root -p -e "CREATE DATABASE "test_task"

dropdb:
	docker exec -it zhiznmart mysql -u root -p -e "DROP DATABASE test_task"

migrateup:
	migrate -path internal/db/migration -database "mysql://root:password@tcp(localhost:3306)/test_task" -verbose up

migratedown:
	migrate -path internal/db/migration -database "mysql://root:password@tcp(localhost:3306)/test_task" -verbose down

sqlc:
	sqlc generate

.PHONY: pullmysql runmysql createdb dropdb migrateup migratedown