# BACKEND GOLANG FOR REST API

## Dtabase Migrations & Connection

Untuk mengelola migrasi database pada project ini, kita menggunakan package [golang-migrate](https://github.com/golang-migrate/migrate).

### ⚙️ Install migrate CLI
Jalankan perintah berikut untuk menginstall migrate CLI:
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` 
`migrate create -ext sql -dir db/migrations create_users_table`

command for migrate and connecting database : 
`migrate -path db/migrations -database "mysql://USERNAME:PASSWORD@tcp(127.0.0.1:3306)/NAMA_DB" up`
command for rollback : 
`migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down`
command for rollback 1 step : 
`migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down 1`

![alt text](image.png)
![alt text](image-1.png)
