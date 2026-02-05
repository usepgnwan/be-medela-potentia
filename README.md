 TEST BE - pt medela potentia tbk.
## Deskripis BE
- Penerapan Swagger Documentation
- Penerapan AUTH JWT
- Penerpan ENV
- Penerepan Pagination
- Buat User Baru
- Data Role / Actor
- Database Postgress
## Access API TEST Swagger 
account
````
username : tes_bemedelapotentia
password : password12345
````

## 1. Setup Apps
- setup fiber & gorm
````
github.com/gofiber/fiber/v2
github.com/joho/godotenv
gorm.io/driver/postgres
gorm.io/gorm
````
- setup validation
````
github.com/go-playground/locales/en
github.com/go-playground/locales/id
github.com/go-playground/universal-translator
github.com/go-playground/validator/v10
github.com/go-playground/validator/v10/translations/id
````
- setup swagger
````
github.com/gofiber/swagger

swag init
```` 
## 2. End Point
- [GET]     api/roles :  api untuk actor / hak access
- [POST]    api/roles :  api add actor / hak access
- [GET]     api/users :  api user list 
- [POST]    api/users :  api registrasi akun
- [POST]    api/users/login :  api generate jwt access
- [POST]    api/users/login/check-jwt :  buat check jwt dari login
- [GET]     api/workflows : api list workflow
- [GET]     api/workflows/id : api detail workflow
- [POST]    api/workflows : api creat workflow
- [GET]     api/workflows/{id}/steps : api list step workflow list
- [POST]    api/workflows/{id}/steps : api creat workflow-step
- api/documentation/swagger/index.html#/ : api swagger
## 3. Testing
- swagger endpoin
- validation form
 
## 4. Instalasi apps 
- git clone
- go mod tidy 
- copy env.example
- go run main.go

## 5. Alur Penggunaan
- buat account role
- buat user, dan pilih rolenya
- login , gunakan jwt token untuk membuat workflow
- buat workflow
- tambah step workflow
- end point request akan menjadi approval pertama & akan   cek ke step workflow untuk ammount / accessnya jika sesuai ambil id requestnya
- lakukan reject masukan id request, akan di cek ke step workflow ammount / accessnya
- lakukan approval masukan id request, akan di cek ke step workflow ammount / accessnya

## 6. example database 
- ada di root dengan nama db_bemedelapotentia_db_bemedelapotentia-202602051348
- account : 
```
staff 
username : staff
password : password
accounting
username : accounting
password : password
manager
username : manager
password : password
```

## Lampiran 

<img width="1919" height="969" alt="image" src="https://github.com/user-attachments/assets/5f5a22a5-f391-48c5-9db7-6ccb9e9ca5b4" />
<img width="1919" height="974" alt="image" src="https://github.com/user-attachments/assets/76a9e468-b263-4453-82a5-d46651542047" />

