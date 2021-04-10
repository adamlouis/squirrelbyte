# squirrelbyte 🐿️

`squirrelbyte` is a "proof of concept" document / search server backed by sqlite.

see it at [https://squirrelbyte.com/](https://squirrelbyte.com/)

# server

`server` is the squirrelbyte server

it's a json blob server

JSON documents are stashed in sqlite using the sqlite json1 extension.

it supports a query syntax similar to [jsonlogic](https://jsonlogic.com/)

it uses sqlite as a data store

the `/documents:search` endpoint DOES NOT use prepared queries, so it's recommended `/documents:search` is used only in safe contexts (read-only, local)

feels verbose ... too much indirection ... but the show must go on

it serves the static web bundle

perhaps more to come ... getting it out there

# web

`web` is the squirrel byte web application.

`web` is written in React is js & with frequent use of the `lodash` library.

the project was started using the `create-react-app` utility.

it uses to `react-virtualized` to maintain good performance queries return many results.

few other deps

the code quality & aesthetics are rough ... performance & functionality should be good.

perhaps more to come!


## why

I like to explore & understand data from the software services I use - Strava, Garmin, GitHub, AWS, & some others.

Some tools I like for exploring data in general are:
- Datasette - for exploring sqlite databases
- jq - for sifting through local json files
- Honeycomb - for general observability of distributed systems ... but in this case, for the query UI & how it works nicely for high cardinality data.
- Grafana / ElasticSearch + Kibana - for general dashboard building, data ingestion, etc.

For my usecase, I wanted a way to:
- Stash my data in its "original" JSON form
- Explore it later & build whatever views I want
- Keep costs & infrastructure complexity low
- Self-host it / own my data

`squirrelbyte` is first step towards these goals -- a document / search server & UI, drawing inspiration from tools that I like. What's here is minimal, but could become more.
