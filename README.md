# Mini Commerce API

This is API for simple E-Commerce

- [Features](#features)
- [Usage](#usage)
- [Further Development](#further-development)

## Features

- Provide Restful API (GET, POST, PUT, DELETE)
- Using Relational Database (in this case, i use many-to-many relationship)
- Primitive SQL Query
- MVC Pattern (I'm using this pattern because good for maintain monolith service and used by many people)
- Redis for cache (set and get key when use method `GET`, unset all key when using method `POST`, `PUT`, and `DELETE`)

## Usage

- Running the API

`make run`

- Rebuild app

`make update-app`

- Rebuild db

`make update-db`

- Stop the API

`make down`

## Further Development

- Using ORM for connection db
- Implement unit test and integration test
- Implement CLEAN code for this project
