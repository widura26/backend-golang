# BACKEND GOLANG FOR REST API

## Dtabase Migrations & Connection

Untuk mengelola migrasi database pada project ini, kita menggunakan package [golang-migrate](https://github.com/golang-migrate/migrate).

### ⚙️ Install migrate CLI
<span>Jalankan perintah berikut untuk menginstall migration CLI:</span>

```go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest``` 

<span>Lalu jalankan perintah untuk membuat file migration. Pada contoh tersebut membuat tabel users</span>

```migrate create -ext sql -dir db/migrations create_users_table```

<span>command for migrate and connecting database :</span>

`migrate -path db/migrations -database "mysql://USERNAME:PASSWORD@tcp(127.0.0.1:3306)/NAMA_DB" up`


<span>command for rollback : </span>

`migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down`

<span>command for rollback 1 step : </span>

`migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down 1`

![alt text](image.png)
![alt text](image-1.png)
