# Blockaction RESTful API

## Environment Variables

* DEBUG: bool, [true|false]
* LOG_LEVEL: string, [ERROR|WARNING|INFO|DEBUG]
* DATABASE_URL: string, ref: [DSN](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)

## Container

### Build container

``` shell
docker build -t blockaction-api:latest .
```

### Run container

``` shell
docker run -itd \
  --name blockaction_api \
  -p 8080:8080 \
  -e LOG_LEVEL=WARNING \
  -e DATABASE_URL="host=127.0.0.1 user=user password=password dbname=blockaction port=5432 sslmode=disable TimeZone=Asia/Taipei" \
  blockaction-api:latest

```

### Endpoints

* POST /auth/signin
* POST /auth/signout
* POST /auth/signup
* GET /api/v1/users
* POST /api/v1/users
* GET /api/v1/users/:username
