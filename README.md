# snack-api

A toy project aimed at learning people's favorite snacks.

## About

`snack-api` is a RESTful API developed in Go.

## Building

### Prerequisites

[golang](https://go.dev/doc/install)

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

### Running

In the terminal, run:

```
snack-api
```

A server will spin up on `http://localhost:8000`.
