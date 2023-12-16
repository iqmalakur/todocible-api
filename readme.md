# ToDocible REST API

This is the REST API for the ToDocible application.

ToDocible is a simple application for todolist management.

You can use the ToDocible REST API by accessing the url `http://api.todocible-akur.koyeb.app/todos/` or `http://localhost:8000/todos/` if you install this project locally.

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
go build -o ToDocible
./ToDocible
```

If you are a Windows user, you can use the command below

```bash
go build -o ToDocible.exe
ToDocible
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
      "id": "e883ab98-865d-4e6d-beb4-bf022ed77f05",
      "title": "Lorem",
      "description": "Lorem Ipsum Dolor Sit Amet.",
      "due_date": "2023-12-16T09:18:17.287Z",
      "completed": false
    },
    {
      "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
      "title": "Consectetur",
      "description": "Sed consectetur accumsan metus.",
      "due_date": "2023-12-16T10:30:00.287Z",
      "completed": true
    }
  ]
}
```

### Get a specific Todo by id

#### Request

`GET /todos/{id}`

Example Request

`GET /todos/45f49812-f8c7-4c2b-8789-8848632d4fac`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success get todo",
  "data": {
    "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "due_date": "2023-12-16T10:30:00.287Z",
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
  "description": "Lorem Ipsum Dolor Sit Amet.",
  "due_date": "2023-12-16T09:18:17.287Z"
}
```

#### Response

Example Response

```json
{
  "success": true,
  "message": "success create new todo",
  "data": {
    "id": "e883ab98-865d-4e6d-beb4-bf022ed77f05",
    "title": "Lorem",
    "description": "Lorem Ipsum Dolor Sit Amet.",
    "due_date": "2023-12-16T09:18:17.287Z",
    "completed": false
  }
}
```

### Change a Todo

#### Request

`PUT /todos/{id}`

Example Request

`PUT /todos/45f49812-f8c7-4c2b-8789-8848632d4fac`

##### Request Body

Example Request Body

```json
{
  "title": "Accumsan",
  "description": "Sed consectetur accumsan metus.",
  "due_date": "2023-12-16T10:30:00.287Z"
}
```

#### Response

Example Response

```json
{
  "success": true,
  "message": "success update todo",
  "data": {
    "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
    "title": "Accumsan",
    "description": "Sed consectetur accumsan metus.",
    "due_date": "2023-12-16T10:30:00.287Z",
    "completed": true
  }
}
```

### Delete a Todo

#### Request

`DELETE /todos/{id}`

Example Request

`DELETE /todos/45f49812-f8c7-4c2b-8789-8848632d4fac`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success delete todo",
  "data": {
    "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "due_date": "2023-12-16T10:30:00.287Z",
    "completed": true
  }
}
```

### Change a Todo's completed to done

#### Request

`PUT /todos/{id}/done`

Example Request

`PUT /todos/45f49812-f8c7-4c2b-8789-8848632d4fac/done`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success set done todo",
  "data": {
    "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "due_date": "2023-12-16T10:30:00.287Z",
    "completed": true
  }
}
```

### Change a Todo's completed to undone

#### Request

`PUT /todos/{id}/undone`

Example Request

`PUT /todos/45f49812-f8c7-4c2b-8789-8848632d4fac/undone`

#### Response

Example Response

```json
{
  "success": true,
  "message": "success set undone todo",
  "data": {
    "id": "45f49812-f8c7-4c2b-8789-8848632d4fac",
    "title": "Consectetur",
    "description": "Sed consectetur accumsan metus.",
    "due_date": "2023-12-16T10:30:00.287Z",
    "completed": false
  }
}
```
