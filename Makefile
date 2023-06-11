
TABLE_NAME = default
VERSION = 1

create_table:
	migrate create -ext sql -dir database/migration/ -seq ${TABLE_NAME}

migration_up: 
	migrate -database "mysql://root:password@tcp(localhost:3306)/pasarmalamkebaya" -path database/migration up

migration_down: 
	migrate -database "mysql://root:password@tcp(localhost:3306)/pasarmalamkebaya" -path database/migration down

migration_fix: 
	migrate -database "mysql://root:password@tcp(localhost:3306)/pasarmalamkebaya" -path database/migration force ${VERSION}
