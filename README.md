# GO API boilerplate

Exemplo de arquitetura de API utilizando a linguagem Go

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
│   │   └── user.go
│   ├── auth
│   │   ├── api.go
│   │   └── login.go
│   └── client
│       ├── api.go
│       └── user.go
├── main.go
├── middleware
│   ├── cors.go
│   ├── logger.go
│   ├── max_client.go
└── utils
    └── handler.go

```

## Endpoints

### Admin

| Description | http | path |
|:--:|:--:|:--|
| list | GET | /api/admin/user |
| get  | GET | /api/admin/user/:id |
| store | POST | /api/admin/user  |
| update | PUT | /api/admin/user/:id |
| delete | DELETE | /api/admin/user/:id |

### Public

| Description | http | path |
|:--:|:--:|:--|
| list | GET | /api/admin/user |
| get  | GET | /api/admin/user/:id |


### Auth

| Description | http | path |
|:--:|:--:|:--|
| admin | POST | /api/auth/admin/signin |
| public  | POST | /api/auth/client/signin |

## Insomnia

Lista de todas requisições para usar com Insomnia, basta importar :) => [link](./.github/Go-API-boilerplate_2020-06-26.json)

## Build & Deploy

### Standalone

* Build

```bash
make prod
```

* Deploy

```bash
./bin/go-api-boilerplate
```

### Docker

* Build

```bash
make docker
```

* Deploy

```bash
docker run -it --name go-api-boilerplate \
    -p 3000:3000 \
    douglaszuqueto/go-api-boilerplate:latest
```

## Referências
