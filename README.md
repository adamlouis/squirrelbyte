[![server](https://github.com/adamlouis/squirrelbyte/actions/workflows/server.yml/badge.svg)](https://github.com/adamlouis/squirrelbyte/actions/workflows/server.yml)
[![web](https://github.com/adamlouis/squirrelbyte/actions/workflows/web.yml/badge.svg)](https://github.com/adamlouis/squirrelbyte/actions/workflows/web.yml)
[![CodeQL](https://github.com/adamlouis/squirrelbyte/actions/workflows/codeql.yml/badge.svg)](https://github.com/adamlouis/squirrelbyte/actions/workflows/codeql.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamlouis/squirrelbyte)](https://goreportcard.com/report/github.com/adamlouis/squirrelbyte)


# squirrelbyte 🐿️

`squirrelbyte` is a "proof of concept" document / search server backed by sqlite.

See a demo at [https://squirrelbyte.com/](https://squirrelbyte.com/).
See a discussion on [Hacker News](https://news.ycombinator.com/item?id=26766557).

## how to run

Run `squirrelbyte` locally in development mode with `docker-compose up`.

TODO: more here

## code overview

There are 2 main components of `squirrelbyte` represented as top-level directories in the project: `server` and `web`

## server

`server` is the the squirrelbyte JSON / REST API.

It stores JSON documents in SQLite using the SQLite json1 extension & serves them via a REST API.

The codebase attempts to adhere to the [Google API Design Guide](https://cloud.google.com/apis/design) and [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

The API is defined as:

```
GET    /status
GET    /documents
POST   /documents
GET    /documents/{documentID}
PUT    /documents/{documentID}
DELETE /documents/{documentID}
POST   /documents:search
```

Where a `document` resource is:

```
{
    "id": "some-unique-identifier",
    "body" { ... some arbitrary json object ... },
    "header" { ... some arbitrary json object ... },
    "created_at": "2021-04-11T21:22:53",
    "updated_at":  "2021-04-11T21:22:53",
}
```

The search endpoint, `POST /documents:search`, supports a query sytax based on 1) SQL and 2) [jsonlogic](https://jsonlogic.com/). The POST body takes the form:

```
{
  "select": [ ... json logic expressions ... ],
  "where": json logic expression,
  "group_by": [ ... json logic expressions ... ],
  "order_by": [ ... json logic expressions ... ],
  "limit": 1000
}
```

See the [jsonlogic](https://jsonlogic.com/) documentation for details on the syntax. Note that the `squirrelbyte` supported operators are different than those in the JSON logic docs. See the `squirrelbyte` supported operators [in the code here](https://github.com/adamlouis/squirrelbyte/blob/main/server/internal/pkg/document/jsonlogic/jsonlogic.go#L21).

In production mode, `server` also serves static web assets.

In development mode, `web` and `server` are run as separate processes and `web` proxies requests to `server`.

## web

`web` is the `squirrelbyte` web application.

`web` is written in React JS.

The project was creating the using the `create-react-app` utility.

`web` is the frontend UI for `server`.

## why

I like to explore & understand data from the software services I use - Strava, Garmin, GitHub, AWS, & some others.

Some tools I like for exploring data in general are:
- [Datasette](https://datasette.io/) - for exploring sqlite databases
- [jq](https://stedolan.github.io/jq/) - for sifting through local json files
- [Honeycomb](https://www.honeycomb.io/overview/) - for general observability of distributed systems ... but in this case, for the query UI & how it works nicely for high cardinality data.
- [Grafana](https://grafana.com/) / [ElasticSearch + Kibana](https://www.elastic.co/demos) - for general dashboard building, data ingestion, etc.

For my usecase, I wanted a way to:
- Stash my data in its "original" JSON form
- Explore it later & build whatever views I want
- Keep costs & infrastructure complexity low
- Self-host it / own my data

`squirrelbyte` is first step towards these goals -- a document / search server & UI, drawing inspiration from tools that I like.

What's here is minimal, but could become more.
