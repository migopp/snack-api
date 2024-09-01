# snack-api

A toy project aimed at learning people's favorite snacks.

## About

`snack-api` is a RESTful API developed in Go. It's currently a local build that leverages PostgreSQL for data storage.

### Snacker Schema

Each profile is attached to a `snacker` data structure:

```json
{
  "id": 1,
  "firstName": "Jane",
  "lastName": "Doe",
  "snack": "Chef Hong Liangpi",
  "hearts": 15
}
```

### Endpoints

- `POST /users`: Creates new user with params `firstName`, `lastName`, and `snack`
- `GET /users`: Gets all users
- `PUT /users/{id}`: Updates the user `{id}`
- `GET /users/{id}`: Gets the user with `{id}`
- `DELETE /users/{id}`: Deletes user with `{id}`
- `PUT /heart/{id}`: Adds a heart to user with `{id}`

## Building

### Prerequisites

1. [golang](https://go.dev/doc/install)
2. [docker](https://www.docker.com/)

### Install

Clone this repo:

```
git@github.com:migopp/snack-api.git
```

Then build:

```
cd snack-api/cmd/snack-api && go install
```

This should install the binary to your `~/go/bin`. Add this directory to your `PATH` if it's not already there.

You'll also need to spin up a PostgreSQL DB.

```
docker run --name pg-test -e POSTGRES_PASSWORD=test -p 5432:5432 -d postgres
docker exec pg-test createdb -U postgres snackapi
```

Keep the settings the same as the commands above.

### Running

In the terminal, run:

```
snack-api
```

A server will spin up on `http://localhost:8000`.
