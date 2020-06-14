# go-api-boilerplate
Exemplo de API utilizando a linguagem Go

## Introdução

## Arquitetura

* 3 áreas
    * admin
    * client
    * auth

* Storage
    * Memory

```bash
http
├── api
│   ├── admin
│   │   ├── api.go
│   │   ├── client.go
│   │   ├── device.go
│   │   ├── farm.go											
│   │   ├── farm_user.go
│   │   └── user.go
│   ├── auth
│   │   ├── api.go
│   │   └── login.go
│   └── client
│       ├── api.go
│       ├── farm.go
│       └── user.go
├── main.go
├── middleware
│   ├── acl.go
│   ├── cors.go
│   ├── jwt.go
│   ├── logger.go
│   ├── maxClient.go
│   ├── rateLimit.go
│   └── validators.go
└── utils
    └── handler.go

```

## Referências