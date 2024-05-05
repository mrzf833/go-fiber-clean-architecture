isi dari docs

- spec api
- isi sql default
- postman
- cara penggunaan api
- cara penggunaan docker



disini saya mencoba membuat go clean architecture dengan menggunakan :
- golang -> require go 1.19 keatas
- mysql -> selesai ✅
- docker
- postman -> sambil update (sudah diterapkan)
- swagger -> sudah package
- jwt -> sudah package ✅
- redis -> sudah package ✅
- penerapan queue
- penerapan cli -> sudah package
- elastic search -> belum
- monitoring queue
- websocket -> belum
- handle file/image dan menambahkan validation -> selesai ✅

package yang digunakan :
- go get github.com/stretchr/testify
- go get github.com/go-playground/validator/v10
- go get github.com/gofiber/fiber/v2
- go get -u gorm.io/gorm
- - go get -u gorm.io/driver/sqlite
- - go get -u gorm.io/driver/mysql
- go get github.com/redis/go-redis/v9
- go get -u github.com/gofiber/contrib/jwt
- go get -u github.com/golang-jwt/jwt/v5
- go get github.com/gofiber/contrib/swagger
- go install github.com/gofiber/cli/fiber@latest
- go get github.com/joho/godotenv
-  go get github.com/DATA-DOG/go-sqlmock

matiin action save di goland

ter inspirasi dari : https://github.com/bxcodec/go-clean-arch yang branch v3
