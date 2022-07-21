# User-API

##### Sebuah service API untuk memanage user
User dibagi menjadi 2 role, **User** dan **Admin**
**Untuk membuat User dengan role Admin, gunakan endpoint** ```localhost:8080/admin```
Input yang dibutuhkan untuk membuat user Admin: **name, email, password**.

## Fitur
#### 1.  Menampilkan data seluruh User
Hak akses: **User & Admin**
Method: **GET**
endpoint ```localhost:8080/users```
Fitur ini dilengkapi dengan fungsi:
- **Filter** berdasarkan Role user (contoh endpoint ```localhost:8080/users?role=user```)
- **Search** berdasarkan Nama user (contoh endpoint ```localhost:8080/users?name=fajar```)
- **Sort** secara Ascending/Descending (contoh endpoint ```localhost:8080/users?sort=asc```)
- **Pagination** dengan limit 5 user yang ditampilkan (contoh endpoint ```localhost:8080/users?page=2```)
 
> Note: Fungsi **Filter**, **Search**, **Sort**, dan **Pagination** bisa digunakan bersamaan

#### 2.  Menampilkan data satu User (User tertentu)
Hak akses: **User & Admin**
Method: **GET**
endpoint ```localhost:8080/users/{id}```

#### 3. Menampilkan data pribadi (User profile)
Hak akses: **User & Admin**
Method: **GET**
endpoint ```localhost:8080/profile```

#### 4. Membuat User baru
Hak akses:  **Admin**
Method: **POST**
Input yang dibutuhkan: **name, email, password**
endpoint ```localhost:8080/users```
> Note: User yang dibuat memiliki default role "user"

#### 5. Mengubah data User
Hak akses:  **Admin**
Method: **PATCH**
Input yang dibutuhkan: **name, email, password**
endpoint ```localhost:8080/users/{id}```


#### 6. Menghapus User
Hak akses:  **Admin**
Method: **DELETE**
endpoint ```localhost:8080/users/{id}```


## Setup
#### - Database
Buat database kosong dengan nama **db_user** seuai pada file ```/config/db.go```
```sh
username := "root"
password := "root"
host := "tcp(127.0.0.1:3306)"
database := "db_user"
```
> Note: Pastikan pengaturan koneksi ke databasenya sesuai dengan **environment** milik anda

#### - Library
Install beberapa library yang dibutuhkan.
```sh
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/golang-jwt/jwt
go get golang.org/x/crypto 
go get github.com/joho/godotenv
```


