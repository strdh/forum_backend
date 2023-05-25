## Introduction
This is a REST API that has functionality as forum app backend,
I created this REST API as a portfolio project to demonstrate my skills in Go, MySQL, and JWT. I learned a lot about these technologies by working on this project, and I hope that it will help me to land a job as a software engineer.

## Tech stack
Tech stack that I used for build this is

[![Go][Go]][Go_URL] [![MYSQL][MYSQL]][MYSQL_URL] [![JWT][JWT]][JWT_URL]
<ul>
  <li>Go programming language</li>
  <li>MYSQL</li>
  <li>JWT</li>
 </ul>
 
 Lib that I used on this project
 
* [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
* [github.com/google/uuid](https://github.com/google/uuid)
* [github.com/gorilla/mux](https://github.com/gorilla/mux)
* [github.com/joho/godotenv](https://github.com/joho/godotenv)
* [github.com/stretchr/testify](https://github.com/stretchr/testify)
* [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)

[Go]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go_URL]: https://go.dev/
[MYSQL]: https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white
[MYSQL_URL]: https://www.mysql.com/
[JWT]: https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens
[JWT_URL]: https://jwt.io/

[driver]: https://github.com/go-sql-driver/mysql


## Features
#### Auth
<ul>
  <li>Register</li>
  <li>Login</li>
</ul>

```sh
  router.HandleFunc("/register", authHandler.Register).Methods("POST")
  router.HandleFunc("/login", authHandler.Login).Methods("POST")
 ```

#### Forum
<ul>
  <li>Get All Forum</li>
  <li>Create Forum</li>
  <li>Get Forum by Id</li>
  <li>Update Forum</li>
  <li>Delete Forum</li>
  <li>Search Forum</li>
  <li>Report Forum</li>
</ul>

```sh
    router.HandleFunc("/forums", forumHandler.Forums).Methods("GET")
    router.HandleFunc("/forums", middleware.AuthMiddleware(forumHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id}", forumHandler.ById).Methods("GET")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Delete)).Methods("DELETE")
    router.HandleFunc("/forums/search", forumHandler.FindForum).Methods("POST")
    router.HandleFunc("/forums/{id}/report", middleware.AuthMiddleware(reportForumHandler.Create)).Methods("POST")
```

#### Message
<ul>
  <li>Read Message(see all messages on a forum)</li>
  <li>Create Message(send a message on a forum)</li>
  <li>Update Message</li>
  <li>Delete Message</li>
  <li>Report Message</li>
  <li>Search Message(search a message on a forum)</li>
</ul>

```sh
    router.HandleFunc("/forums/{id_forum}/messages", messageHandler.ByIdForum).Methods("GET")
    router.HandleFunc("/forums/{id_forum}/messages", middleware.AuthMiddleware(messageHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Delete)).Methods("DELETE")
    router.HandleFunc("/forums/{id_forum}/messages/{id}/report", middleware.AuthMiddleware(reportMessageHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id}/searchmsg", forumHandler.FindMsg).Methods("POST")
```

## Database Schema
Database on this REST API consists of 7 tables
![My Image](db_schema.png)

## How to Use 
* Download this repo or clone it by using `git clone https://github.com/strdh/forum_backend.git`
* Create a database on mysql or use can use mariadb
* Crate tables based on the schema or you can just import the sql file that provided in this repo
* Create `.env` file on the root directory and you can see `.env.example`
* adjust the the enviroment value in `.env` based on value on your system
  * `ADDRESS=your_address` example `www.example.com`
  * `JWT_KEY=your_jwtkey` this key is used for generate a JWT and verify that
  * `DB_HOST=your_dbhost` example `localhost`
  * `DB_PORT=your_dbport` example `3306`
  * `DB_USERNAME=your_dbusername` example `root`
  * `DB_PASSWORD=your_dbpassword` example `secret-password***&(*&*"`
  * `DB_NAME=your_dbname` example `forum_db`
  * `DB_TEST_NAME=your_dbtestname` this is needed if you wanna separate db test
* run with `go run main.go`

## Detail
#### Dir structure

```sh
    forum_backend
    |
    |__config
    |_____database.go
    |__handlers
    |_____auth_handler.go
    |_____forum_handler.go
    |_____message_handler.go
    |_____report_forum_handler.go
    |_____report_message_handler.go
    |__models
    |_____forum.go
    |_____message.go
    |_____report_forum.go
    |_____report_message.go
    |_____topic.go
    |_____user.go
    |__test
    |_____test1.go
    |_____test2.go
    |_____testn.go
    |__utils
    |_____jwt.go
    |_____response.go
    |__validators
    |_____auth.go
    |_____forum.go
    |_____message.go
    |_____report_forum.go
    |_____report_message.go
    |__.env
    |__.env.example
    |__.gitignore
    |__apispecs
    |__go.mod
    |__go.sum
    |__main.go

```

  
