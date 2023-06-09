# Full Cycle Learning Experience

## Chat Service
* GoLang

### Definitions
* **Domain** Core of application. Business rules
* **Gateway** Contracts defineds by domain to access external. Works like interfaces in OOP
* **UseCase** Intenção do usuário
* **RRPC** Faster then Rest. Work with stream

### 1 - Commands ChatService:
```Bash
$ go mod init github.com/rlinsdev/FC-LX/tree/main/ChatService,
# To install github.com/google/uuid
$ /ChatService> go mod tidy
# To install tiktoken_go "github.com/j178/tiktoken-go" 
$ /ChatService> go mod tidy
```
### 2 Run:
```Bash
# To run:
$ go run cmd/chatservsice/main.go 
```

### 2 - Commands migration:
```Bash
# Install migration on go
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# Migration to MySql 
$ /ChatService> migrate create -ext=mysql -dir=sql/migrations -seq init
# Generate files in infra/db
$ /ChatService> sqlc generate
```

### 3 NestJs:
```Bash
$ npx create-next-app --typescript 
```

## Links Docs
* [OpenAI API docs](https://platform.openai.com/docs/api-reference/introduction)
* [OpenAI](https://openai.com/)
* [SQLC.dev](https://sqlc.dev/)
* [Protocol Buffers](https://protobuf.dev/)


## Links Course
* [FullCycle Page](https://fcexperience.fullcycle.com.br/evento/)
* [FullCycle YouTube - 1](https://www.youtube.com/watch?v=UugkE-OeE4E)
* [FullCycle YouTube - 2](https://www.youtube.com/watch?v=lstRv2q-sOI)
* [Repo URL](https://github.com/devfullcycle/fclx)


