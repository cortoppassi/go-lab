# Go Backend Lab

Backend simples em Go com CRUD de tarefas, organizado em camadas.

## Estrutura

- `cmd/api`: ponto de entrada da aplicacao
- `internal/httpapi`: roteamento e helpers HTTP
- `internal/task`: dominio de tarefas com handler, service e repository

## Como rodar

```bash
go run ./cmd/api
```
