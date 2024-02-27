# golang-gin-todos_api
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/gin_gonic_logo.png)
# organizer

| nickname | full-name           | github username |
| -------- | ------------------- | --------------- |
| Tiger    | Korrakot Triwichian | korrakot1396    |

## Project installation process

_For server_

```shell
cd golang-gin-todos
go mod download
go mod tidy
go run cmd/main.go
```

_For unit test_

```shell
cd golang-gin-todos
cd ./tests
go test
```

_For postman_

```shell
open postman application
cd postman
export Todo-App-Api.postman_collection.json into your-postman-application
```
## Swagger Documentation

_For swagger_

```shell
cd golang-gin-todos
go run cmd/main.go
open brower go to "http://localhost:8081/docs/index.html"
```

## Swagger
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_1.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_2.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_3.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_4.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_5.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/swagger_6.png)

## Functionality

1. CRUD todos for view, create, update, delete
2. Feature for create todos
3. Validation 100 Character for Field 'Title'
4. Use UUID supported for Field 'ID'
5. Use Gingonic Framework in Golang
6. Create Unit Testing in Golang
7. Validation Duplicate Field 'Title'



## Demo Server Run in Postman And Unit Test

![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/unit_testing.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/get_all_tasks.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/add_todo.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/duplicate.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/update_task.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/mark_done.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/delete_by_id.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/delete_all.png)
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/pv_1.png)
