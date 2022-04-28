


# Web Based ToDo List Application

![Uygulama Ekran Görüntüsü](https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/goroutines/png/2.png)

- [x] 1- User interface for ONLY adding ToDo’s
- [x] 2- Back-end service to store a persistent state of ToDo list
- [x] 3- Writing deployment files of your front-end and back-end
- [x] 4- Automating build, test and deployment of your application via CI/CD pipeline

## Tech Stack

**Tech Stack:**  Golang, stdlib , In-memory, Pact-go
## Directory Structure
```
.backend-todo
├── .config
│   ├── local.json  
├── config 
│   ├── config.go  
│   ├── env.go  
├── handler 
│   ├── handler.go  
│   ├── handler_test.go  
├── mocks
│   ├── mock_repository.go  
│   ├── mock_service.go  
├── model
│   ├── model.go  
├── pact   
│   ├── log
│       ├── pact.log
│   ├── provider_test.go  
├── repository 
│   ├── repository.go
├── server
│   ├── server.go
├── service
│   ├── service.go
│   ├── service_test.go   
├── .gitlab-ci.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go

```


## Model

This project contains in-memory db.

Todo model:

```bash 
  type Todo struct {
	ID   int       `json:"id"`
	Todo string    `json:"todo"`
	Time time.Time `json:"time"`
}
```

## Repository

In this project, I used interfaces for each package to write tests.

```bash 
type IRepositoryTodo interface {
	GetTodos() map[string]*model.Todo
	AddTodo(todo model.Todo) *model.Todo
}
```
After using Interface, I used the [github.com/golang/mock](https://github.com/golang/mock) package to mock these packages.

Mock command:

`mockgen -source service/service.go -destination mocks/mock_repository.go -package mocks IRepositoryTodo`


![Uygulama Ekran Görüntüsü](https://3903010379-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-L9Tqx5WSaiE4u24Pk05-2910905616%2Fuploads%2Fgit-blob-cad524fa8cb34476d131615dfd4861f9aa63a7c4%2Fred-green-blue-gophers-smaller.png?alt=media)



## Test
You can call the test files in the project with this command:

`go test ./... -v`

or file:

`go test ./handler/handler_test.go`

## Server
To run the application's server:

`go run .`

Development server : [localhost:3000]()




## API

#### Get all todo list

`GET /api/v1/todos`

```http
  curl localhost:3000/api/v1/todos
```


#### Create new todo item

`POST /api/v1/todos`

```http
  curl -H "Content-type: application/json" -X POST -d '{"todo":"dummy"}' localhost:3000/api/v1/todos
```



#### Delete all todo list

` DELETE /api/v1/todos/deleteAll`

```http
  curl -X DELETE localhost:3000/api/v1/todos/deleteAll
 
```
 




## Licenses

[MIT](https://choosealicense.com/licenses/mit/)

  
