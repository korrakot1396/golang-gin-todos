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

1. Comprehensive Todo Management: Implement full CRUD functionality to facilitate the viewing, creation, updating, and deletion of todos.
2. Streamlined Todo Creation: Develop a user-friendly feature for seamlessly creating new todos.
3. Title Field Validation: Implement validation to restrict the title field to a maximum of 100 characters, ensuring data integrity and user experience.
4. Utilize UUID for ID Generation: Utilize universally unique identifiers (UUID) to generate IDs for todos, ensuring uniqueness and reliability.
5. Integration with Gingonic Framework: Harness the power of the Gingonic framework in Golang for robust backend development, ensuring scalability and performance.
6. Unit Testing for Code Quality: Implement comprehensive unit tests in Golang to verify the functionality and quality of the codebase.
7. Prevent Duplicate Titles: Incorporate validation to prevent the creation of todos with duplicate titles, enhancing data consistency and usability.
8. Efficient Image Upload and Conversion: Enable users to upload images for their todos and seamlessly convert them to base64 format, enhancing data handling efficiency.
9. Flexible Sorting Options: Provide users with the flexibility to sort todos based on title, status, and date, enhancing usability and navigation.
10. API Exploration with Swagger: Develop an API explorer using Swagger to facilitate seamless exploration and interaction with the API endpoints, enhancing developer experience and productivity.



## Demo Server Run in Postman And Unit Test

![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/unit_testing.png)
### Get all task
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/get_all_tasks.png)
### Add new task
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/add_todo.png)
### Case Title duplicated
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/duplicate.png)
### Update task
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/update_task.png)
### Mark done from 'IN_PROGRESS' to 'SUCCESS'
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/mark_done.png)
### Delete by Id
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/delete_by_id.png)
### Delete all tasks
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/delete_all.png)
### Save Base64 into Field 'image'
![](https://github.com/korrakot1396/golang-gin-todos/blob/main/img/pv_1.png)
