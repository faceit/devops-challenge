# Ops Demo App

This application is a web api written in golang which connects to a PostgreSQL database and runs PING command against the database.

The application looks for the following environment variables:

- `PORT` - the http port to start the health endpoint on, by default `8080`
- `POSTGRESQL_HOST` - `localhost`
- `POSTGRESQL_PORT` - `5432`
- `POSTGRESQL_USER` - `postgres`
- `POSTGRESQL_PASSWORD` - `mysecretpassword`
- `POSTGRESQL_DBNAME` - `postgres`

If the application fails to start for whatever reason it will typically give a useful error message which indicates the problem its facing.

When the application starts it attempts to connect to postgres. If it cannot do this it will fail to start and the application will exit with a error message.

This application has a single endpoint `/health`. The `/health` endpoint attempts to run a PING command against postgres.

A successful response looks like:

```bash
➜ curl -v http://localhost:8080/health
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /health HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.74.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Wed, 09 Feb 2022 15:21:30 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
< 
UP
```

An unsuccessful response is the following:

```bash
➜ curl -v http://localhost:8080/health
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /health HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.74.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 500 Internal Server Error
< Date: Wed, 09 Feb 2022 15:22:48 GMT
< Content-Length: 5
< Content-Type: text/plain; charset=utf-8
< 
DOWN
```

## How to build the application

The following command can be used to build the application if running on linux and will produce you a working binary.

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/ops-test-app
```

## How to run the app locally to test

First we need a running postgres server, if you don't already have one it can be achieved using docker:

```bash
docker run --rm --name=postrgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres:13-alpine
```

Then to run the application itself you can run:

```bash
go run cmd/ops-test-app/main.go
```

## Database

You can create a new user/database for the application if you wish, but it is not part of the test and we aren't expecting you to do that. We just want a working connection to the database where the application can successfully ping the database.
