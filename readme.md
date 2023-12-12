# ToDocible REST API

This is the REST API for the ToDocible application.

ToDocible is a simple application for todolist management.

<details>
<summary>Table of Contents</summary>

- [Installation](#installation)
- [Run the app](#run-the-app)
- [REST API](#rest-api)
  - [Get list of Todo](#get-list-of-todo)
  - [Get a specific Todo by id](#get-a-specific-todo-by-id)
  - [Create new Todo](#create-new-todo)
  - [Change a Todo](#change-a-todo)
  - [Delete a Todo](#delete-a-todo)
  - [Change a Todo's completed to done](#change-a-todos-completed-to-done)
  - [Change a Todo's completed to undone](#change-a-todos-completed-to-undone)

</details>

## Installation

You can install this project locally by cloning this github repository using the command below.

```bash
git clone https://github.com/iqmalakur/todocible-api.git
```

## Run the app

```bash
go build -o todolist
./todolist
```

If you are a Windows user, you can use the command below

```bash
go build -o todolist.exe
todolist
```

## REST API

### Get list of Todo

#### Request

`GET /todos/`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success get all todos",
  "data": [
    {
      "id": 0,
      "title": "Lorem",
      "description": "Lorem Ipsum Dolor Sit Amet.",
      "completed": false
    },
    {
      "id": 1,
      "title": "Consectetur",
      "description": "Sed consectetur accumsan metus.",
      "completed": true
    }
  ]
}
```

### Get a specific Todo by id

#### Request

`GET /todos/{id}`

Example Request

`GET /todos/1`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success get todo",
  "data": {
    "id": 1,
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "completed": true
  }
}
```

### Create new Todo

#### Request

`POST /todos/`

##### Request Body

Example Request Body

```json
{
  "title": "Lorem",
  "description": "Lorem Ipsum Dolor Sit Amet."
}
```

#### Response

Example Response

```json
{
  "success": true,
  "message": "success create new todo",
  "data": {
    "id": 0,
    "title": "Lorem",
    "description": "Lorem Ipsum Dolor Sit Amet.",
    "completed": false
  }
}
```

### Change a Todo

#### Request

`PUT /todos/{id}`

Example Request

`PUT /todos/1`

##### Request Body

Example Request Body

```json
{
  "title": "Accumsan",
  "description": "Sed consectetur accumsan metus."
}
```

#### Response

Example Response

```json
{
  "success": true,
  "message": "success update todo",
  "data": {
    "id": 1,
    "title": "Accumsan",
    "description": "Sed consectetur accumsan metus.",
    "completed": true
  }
}
```

### Delete a Todo

#### Request

`DELETE /todos/{id}`

Example Request

`DELETE /todos/1`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success delete todo",
  "data": {
    "id": 1,
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "completed": true
  }
}
```

### Change a Todo's completed to done

#### Request

`PUT /todos/{id}`

Example Request

`PUT /todos/1/done`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success set done todo",
  "data": {
    "id": 1,
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "completed": true
  }
}
```

### Change a Todo's completed to undone

#### Request

`PUT /todos/{id}`

Example Request

`PUT /todos/1/undone`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success set undone todo",
  "data": {
    "id": 1,
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "completed": false
  }
}
```
