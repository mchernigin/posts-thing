ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migration_up:
	migrate -path database/migrations/ -database "${DATABASE_URL}" -verbose up

migration_down:
	migrate -path database/migrations/ -database "${DATABASE_URL}" -verbose down

migration_reset: migration_down migration_up

migration_drop:
	migrate -path database/migrations/ -database "${DATABASE_URL}" -verbose drop

