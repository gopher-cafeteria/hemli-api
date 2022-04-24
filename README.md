---
title: Chi
description: A Chi server for Hemli API
tags:
  - chi
  - golang
  - postgres
---

# Hemli API

## ✨ How to run web server

1. Install dependencies

```
$ go get -d ./...
```

2. navigate to the cmd directory

```
$ cd .\cmd\
```

3. run the application

```
$ go run .
```

## 🗄 Database and Migrations

- Hemli API uses Postgresql as main database source. In order to have an up and running postgres DB make sure you have docker installed on your local machine.Run following command to have postgres container up and running.

<hr>

```
$ docker-compose up -d
```

- Once database is up and running, run migrations to have required tables created.

_Please note_: in order to run sql-migrate and all migration commands, you need to have pre-installed CLI tools. More information and instalation instructions check in <a href="https://github.com/rubenv/sql-migrate"> official documentation</a>.

<hr>

```
$ make migrate-up
```

- similarly you can run following command to rollback all migrations

```
$ make migrate-down
```

- In order to create a new migration file, run following command in your terminal.

```
$ sql-migrate new MY_TABLE
```
