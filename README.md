# FAT Trading BOT

## Library
 - Web Frimework [Gin Gonic](https://github.com/gin-gonic/gin) installed (**version v1.7.7**)
 - A SQL Compiler [sqlc](https://github.com/kyleconroy/sqlc) installed (**version v1.11.0**)
 - Database Migrations [golang-migrate](https://github.com/golang-migrate/migrate) installed (**version v4.15.1**)
 - Management Configuration File [viper](https://github.com/spf13/viper) installed (**version v1.10.1**)

 ## Exchange
 - [Indodax](https://indodax.com/)
    * Api Documentation [Github](https://github.com/btcid/indodax-official-api-docs)

## Indicator trading support
 - RSI
 - MACD
 - MA
 - EMA
 - BBANDS
 - SMA
 - WMA
 - KAMA
 - DEMA
 - MACDEXT
 - TEMA

## Installation
To install FAT, you need to install Go and set your Go workspace first.
To install Gin package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.13+ is required**), then you can use the below Go command to install Gin.

```sh
$ go mod init
```
```sh
$ go mod tidy
```

2. Setup Envronment :

```go
ENV_APP=dev
ADDRESS_APP=0.0.0.0:9091

## DATABASE ENVIRONMENT ##
DATABASE_HOST=
DATABASE_PORT=
DATABASE_USER=
DATABASE_PASS=
DATABASE_NAME=

## EXCHANGE ENVIRONMENT ##
INDODAX_MAIN_DOMAIN=https://indodax.com
INDODAX_PUBLIC_KEY=
INDODAX_PRIVATE_KEY=
INDODAX_CRYPTO_CURRENCY=btcidr 
INDODAX_CRYPTO_TYPE=btc 
INDODAX_PAIR=btc_idr
INDODAX_PERIODE_TIME=1440  
INDODAX_TIME_FRAME=1
```

## Quick start
the configuration the start already define on Makefile

```sh
# Makefile
sqlc:
	sqlc generate

generate_migration:
	migrate create -ext sql -dir ./repository/migrations -seq $(table)

dbup:
	migrate -database "postgres://postgres:postgres@localhost:5432/fat?sslmode=disable" -path ./repository/migrations up

dbdown:
	migrate -database "postgres://postgres:postgres@localhost:5432/fat?sslmode=disable" -path ./repository/migrations down

dbrollback:
	migrate -path ./repository/migrations -database "postgres://postgres:postgres@localhost:5432/fat?sslmode=disable" force $(version)

server_dev:
	go run . --environment=dev

server_stg:
	go run . --environment=staging
    
server_prod:
	go run . --environment=production

```

```sh
# example the following below
# this for generate sql on respository
$ make sqlc
```
```sh
# this for generate sql file migration on directory ./repository/migrations
$ make generate_migration table=table_name
```
```sh
# this for migration db up
$ make dbup
```
```sh
# this for migration db dwn
$ make dbdown
```

```sh
# this for rollback dirty db version
$ make dbrollback version=1
```

```sh
# this for run server with environemnet dev
$ make server_dev
```

```sh
# this for run server with environemnet staging
$ make server_stg
```

```sh
# this for run server with environemnet production
$ make server_prod
```