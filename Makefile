.PHONY: init_db
init_db:
	migrate -path ./migrations -database 'postgres://user:password@localhost:5432/client_info?sslmode=disable' --verbose up
.PHONY: drop_db
drop_db:
	migrate -path ./migrations -database 'postgres://user:password@localhost:5432/client_info?sslmode=disable' --verbose down -all
