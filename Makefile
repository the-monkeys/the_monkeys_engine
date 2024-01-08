include .env
export

proto:
	protoc microservices/the_monkeys_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc microservices/auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc microservices/article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc microservices/user_service/service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc microservices/blogsandposts_service/blog_service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc microservices/file_server/service/pb/*.proto --go_out=. --go-grpc_out=.


proto-gen:
	protoc apis/grpc/**/*.proto --go_out=. --go-grpc_out=.

proto-gen-interservices:
	protoc apis/interservice/**/*.proto --go_out=. --go-grpc_out=.


sql-gen:
	echo "Enter the file's name or description (Node keep it short):"
	@read INPUT_VALUE; \
	migrate create -ext sql -dir schema-migrations -seq $$INPUT_VALUE


# TODO: Enable SSL for psql
migrate-up:
	migrate -path schema-migrations -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose up

migrate-down:
	migrate -path schema-migrations -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose down 1

migrate-force:
	echo "Enter a version:"
	@read INPUT_VALUE; \
	migrate -path schema-migrations -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose force $$INPUT_VALUE

run:
	./build.sh
