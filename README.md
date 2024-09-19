create .env

PORT=8000
DB_URL="host=gaudily-stable-weimaraner.data-1.use1.tembo.io user=postgres password=***askforit*** dbname=postgres port=5432"

CompileDaemon -command="./backend" or go run main.go

To migrate (for example create a table)
```go run ./migrate/migrate.go```
