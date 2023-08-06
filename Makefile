include app.env

ifeq (${NODE_ENV}, dev)
	DB_NAME=${DEV_DB_NAME}
	DB_USER=${DEV_DB_USER}
	DB_PASSWORD=${DEV_DB_PASSWORD}
	DB_HOST=${DEV_DB_HOST}
	DB_PORT=${DEV_DB_PORT}
	SSL_MODE=disable
else
	DB_NAME=${PROD_DB_NAME}
	DB_USER=${PROD_DB_USER}
	DB_PASSWORD=${PROD_DB_PASSWORD}
	DB_HOST=${PROD_DB_HOST}
	DB_PORT=${PROD_DB_PORT}
	SSL_MODE=require
endif

migrate:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrate_create:
	migrate create -ext sql -dir db/migrations -seq ${MIGRATE_NAME}_schema

migration_up:
	migrate -path db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -verbose up

migration_down_all:
	migrate -path db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -verbose down

migration_down_by_id:
	migrate -path db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -verbose down ${VERSION}


migration_fix:
	migrate -path db/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" force ${VERSION}

sqlc:
	sqlc compile
	sqlc generate

