 TEST BE - pt medela potentia tbk.
## Deskripis BE
- Penerapan Swagger Documentation
- Penerapan AUTH JWT
- Penerpan ENV
- Penerepan Pagination
- Buat User Baru
- Data Role / Actor
- Database Postgress
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
- api/roles :  api untuk actor / hak access
- api/users :  api user list 
- api/users :  api registrasi akun
- api/users/login :  api generate jwt access
- api/users/login/check-jwt :  buat check jwt dari login
- api/workflows : api list workflow
- api/workflows/id : api detail workflow
- [post] api/workflows : api creat workflow
- api/workflows-step/id : api step workflow list
- [post] api/workflows-step : api creat workflow-step
- api/documentation/swagger/index.html#/ : api swagger
## 3. Testing
- swagger endpoin
- validation form
 
## 4. Instalasi apps 
- git clone
- go mod init 
- copy env.example
- go run main.go
