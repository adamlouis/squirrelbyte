## about

I like to explore & understand data from the software services I use - Strava, Garmin, GitHub, AWS, & some others.

`squirrelbyte` is a "proof of concept" personal hobby project that lets me do that.

For my usecase, I wanted a way to:

- Stash my data in its "original" JSON form
- Explore it later & build whatever views I want
- Keep costs & infrastructure complexity low
- Self-host it / own my data


The application here is a first step towards these goals -- a document / search server & UI.

It draws inspiration from some of the tools that I like:

- [Datasette](https://datasette.io/) - for exploring sqlite databases
- [jq](https://stedolan.github.io/jq/) - for sifting through local json files
- [Honeycomb](https://www.honeycomb.io/overview/) - for general observability of distributed systems ... but in this case, for the query UI & how it works nicely for high cardinality data.
- [Grafana](https://grafana.com/) / [ElasticSearch + Kibana](https://www.elastic.co/demos) - for general dashboard building, data ingestion, etc.

What's here is minimal, but could become more.

## implementation brief

There are currently 2 components of `squirrelbyte` - a web application and a server.

The web application is written in react and interacts with the server.

The server is written and golang & manages CRUD & querying. It stashes JSON documents in sqlite using the [sqlite json1 extension](https://www.sqlite.org/json1.html) and supports a query syntax similar to [jsonlogic](https://jsonlogic.com/). `jsonlogic` is essentially used as a restricted AST for a SQL query.

## some links

- see the initial Hacker News post: [Show HN: Squirrelbyte – a SQLite-based JSON document server](https://news.ycombinator.com/item?id=26766557)
- see the source on GitHub: [`github.com/adamlouis/squirrelbyte`](https://github.com/adamlouis/squirrelbyte)
