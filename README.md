# BACKEND GOLANG FOR REST API
## Migrations & Connect to Database
for configurating migration, you can install package 
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` and then
run the command below
`migrate create -ext sql -dir db/migrations create_users_table`

command for migrate and connecting database : `migrate -path db/migrations -database "mysql://USERNAME:PASSWORD@tcp(127.0.0.1:3306)/NAMA_DB" up`
command for rollback : `migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down`
command for rollback 1 step : `migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down 1`

![alt text](image.png)
