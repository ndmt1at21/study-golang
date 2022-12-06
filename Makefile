migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "mysql://${MYSQL_ROOT_USER}:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" -verbose up  

migratedown:
	migrate -path db/migration -database "mysql://${MYSQL_ROOT_USER}:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" -verbose down  

gen_pb:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=require_unimplemented_servers=false:pb proto/*.proto

start:
	go run main.go --port 8001

test: 
	go test -v -cover ./...

.PHONY: migrateinit migrateup migratedown sqlc gen_pb start test 