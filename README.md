# BACKEND GOLANG FOR REST API
<p>Proyek ini merupakan dummy project untuk pembelajaran semata. Framework yang digunakan pada proyek ada Golang Echo</p>

## Database Migrations & Connection

Untuk mengelola migrasi database pada project ini, kita menggunakan package [golang-migrate](https://github.com/golang-migrate/migrate).

### ⚙️ Install migrate CLI
<span>Jalankan perintah berikut untuk menginstall migration CLI:</span>

```go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest``` 

<span>Lalu jalankan perintah untuk membuat file migration. Pada contoh tersebut membuat tabel users</span>

```migrate create -ext sql -dir db/migrations create_users_table```

<span>command for migrate and connecting database :</span>

`migrate -path db/migrations -database "mysql://USERNAME:PASSWORD@tcp(127.0.0.1:3306)/NAMA_DB" up`


<span>command for rollback : </span>

`migrate -path db/migrations -database "mysql://root:Widura260503@tcp(127.0.0.1:3307)/go_rest_api" down`

<span>command for rollback 1 step : </span>

`migrate -path db/migrations -database "mysql://root:root123@tcp(127.0.0.1:3306)/golang_db" down 1`

## Bagaimana Jika terdapat perubahan kolom atau tabel pada migration
<p>Untuk Development, mungkin dapat menggunakan perintah rollback dan memodifikasi migration yang tersedia. 
Jika dalam tahap production, dapat menjalankan perintah migration baru untuk memodifikasi kolom atau tabel</p>

## Gorm package for ORM
<span>Jalankan perintah di bawah ini :</span>

```go get -u gorm.io/gorm``` 

<span>Dikarenakan pada proyek ini menggunakan mysql, diperlukan untuk menginstal driver mysql yang juga disediakan oleh GORM</span>

```go get -u gorm.io/driver/mysql```


## Other Source


![alt text](image.png)
![alt text](image-1.png)
