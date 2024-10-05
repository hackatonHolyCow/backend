migration:
ifdef NAME 
	@goose -s -dir ./pkg/postgres/migrations create $(NAME) go
else
	@echo "ERROR: must define a migration name"
endif