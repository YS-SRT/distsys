##### Distsys Project:
分级销售系统 + 功能化模块组合

##### Tech Stack:
Gin, Gorm, Viper, Zap, Casbin, JWT, swagger
PostgreSQL, minio, redis

amis, vue3, bootstrap5

Docker, Docker-compose

##### swag cmd
$ go install github.com/swaggo/swag/cmd/swag@latest

##### swagger init
$ swag init -g ./main.go -o ./docs --parseDependency

##### web ui in dev environment

http://localhost:8090/swagger/index.html

minio local:
http://localhost:9091


Admin url
http://localhost:8090/