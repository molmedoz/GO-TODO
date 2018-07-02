# TODO List
Example of go server. The server expose a rest-api specified in the [instruction](INSTRUCTIONS.MD)

The server has a basic authentication.
The project has 2 dependencies:
* mattn/go-sqlite3 for connecting to the database
* golang.org/x/crypto for encrypting the passwords

## End points

/login

### Login
{POST} /login
#### Success
#### Error

### SingUp
{POST} /signup
#### Success
#### Error

### Logout
{POST} /logout

#### Headers
#### Success
#### Error
### TaskList
{GET} /api/todo
#### Headers
#### Success
#### Error
### Add Task
{POST} /api/todo
#### Headers
#### Success
#### Error
### Update Task

{PUT} /api/todo/:taskId
#### Headers
#### Output
### Delete Task
{DELETE} /api/todo/:taskId

#### Headers
#### Output
