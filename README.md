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

Start a postgres container, following the example command below:

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

Prepare yourself a `config.ini` configuration according to `config.ini.tmpl`. Following is the configuration specification:

|     Name     |  Type   | In  |                 Description                  |
|:------------:|:-------:|:---:|:--------------------------------------------:|
|    WORKER    | integer | APP |      Maximum number of parallel workers      |
|  QUEUE_SIZE  | integer | APP |       Maximum number of parallel queue       |
|   LANGUAGE   | string  | APP |        Targeted programming language         |
|   MIN_STAR   | integer | APP |     Minimum number of stars a repo gains     |
| GITHUB_TOKEN | string  | WEB |    GitHub token to unlock API rate limit     |
|    TRYOUT    | integer | WEB | Maximum number of retrying to request a page |
|     HOST     | string  | DB  |            Database host address             |
|     USER     | string  | DB  |              Database user name              |
|   PASSWORD   | string  | DB  |            Database user password            |
|    DBNAME    | string  | DB  |                Database name                 |
|     PORT     | integer | DB  |                Database port                 |

## :hammer_and_wrench: Build

```bash
$ go build GitHunter
```

## :rocket: Run

To crawl the repositories' metadata:

```bash
$ ./GitHunter crawl
```

To clone the repositories:

```bash
$ ./GitHunter clone
```
