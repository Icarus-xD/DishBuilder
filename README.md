# Zhiznmart Test Task [Golang]

## Run (Locally)
### Prerequisites
- go 1.20
- docker
- make (<i>optional</i>, used to run CLI commands)
- [sqlc](https://github.com/kyleconroy/sqlc) (<i>optional</i>, used to re-generate fully type-safe idiomatic Go code from SQL)

For the first time, run `make` commands or what they contain (see `Makefile`) to run app:
```bash
make pullmysql
make createdb
make migrateup
make runapp
```
For the next times, run `make runapp`.

## Tech Stack
- Golang
- Gin
- Viper
- MySQL
- sqlc

## App Description
App is a dish builder that provides a REST API that consists of a single endpoint:
#### GET /dish?recipe=
For that request need to provide query param <strong>recipe</strong> which represents the string of ingredients types included in the dish.
There is an example of a response to the request with query param <i>"dccii"</i>:
```json
[
	{
“products”: [
	{“type”:”Тесто”,”value”:”Тонкое тесто”},
	{“type”:”Сыр”,”value”:”Моцарелла”},
	{“type”:”Начинка”,”value”:”Ветчина”},
	{“type”:”Начинка”,”value”:”Колбаса”},
],
“price”: 215
},
{
“products”: [
	{“type”:”Тесто”,”value”:”Тонкое тесто”},
	{“type”:”Сыр”,”value”:”Моцарелла”},
	{“type”:”Начинка”,”value”:”Ветчина”},
	{“type”:”Начинка”,”value”:” Грибы”},
],
“price”: 235
},
….
]

```