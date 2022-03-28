# GitHunter

[![Build](https://github.com/ZJU-SEC/GitHunter/actions/workflows/build.yml/badge.svg)](https://github.com/ZJU-SEC/GitHunter/actions/workflows/build.yml)

`GitHunter` is a tiny yet powerful crawler infra to collect OSS projects on GitHub. It queries GitHub search API and persist the data into the Postgres database.

[Check here](doc/README.md) to know what the collected data is.

## :gear: Prerequisite

- Docker
- Golang
- PostgreSQL


### :bulb: Dockerized PostgreSQL

To run a dockerized PostgreSQL, check [this](https://hub.docker.com/_/postgres).

Start a postgres container:

```bash
$ docker run \
  --name postgres -d \
  --restart unless-stopped \
  -e POSTGRES_USER=ZJU-SEC \
  -e POSTGRES_PASSWORD=<YOUR DB PASSWORD> \
  -e POSTGRES_DB=GitHunter \
  -p 5432:5432 postgres
```

## :page_facing_up: Make the Configurations

Prepare yourself a `config.ini` configuration according to `config.ini.tmpl`.

### Parameters

| Name          | Type    | In  |Description|
|:-:            |:-:      |:-:  |:-:        |
| WORKER        | integer | APP | Maximum number of parallel workers |
| QUEUE_SIZE    | integer | APP | Maximum number of tuples to insert into database at a time |
| LANGUAGE      | string  | APP | Targeted programming language |
| MIN_STAR      | integer | APP | Minimum number of stars a repo gains |
| GITHUB_TOKEN  | string  | WEB | Github token enabling usage of github API |
| TRYOUT        | integer | WEB | Maximum number of retrying to request a page |
| HOST          | string  | DB  | Database host address |
| USER          | string  | DB  | Database user name |
| PASSWORD      | string  | DB  | Database user password |
| DBNAME        | string  | DB  | Database name |
| PORT          | integer | DB  | Database port |

## :hammer_and_wrench: Build

```bash
$ go build CIHunter
```

## :rocket: Run

```bash
$ ./CIHunter
```
