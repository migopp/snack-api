# snack-api

A toy project aimed at learning people's favorite snacks.

## About

`snack-api` is a RESTful API developed in Go.

## Building

### Prerequisites

[golang](https://go.dev/doc/install)
[docker](https://www.docker.com/)

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

You'll also need to spin up a PostgreSQL Database.

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
