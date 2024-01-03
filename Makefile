include .env
export

proto:
	protoc services/api_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/user_service/service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/blogsandposts_service/blog_service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/file_server/service/pb/*.proto --go_out=. --go-grpc_out=.


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
