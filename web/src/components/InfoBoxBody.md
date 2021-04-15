# why

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

The web application here is a first step towards these goals -- a document / search server & UI, drawing inspiration from tools that I like.

What's here is minimal, but could become more.

# what

`squirrelbyte` is a "proof of concept" document / search server backed by sqlite. JSON documents are stashed in sqlite using the [sqlite json1 extension](https://www.sqlite.org/json1.html). It supports a query syntax similar to [jsonlogic](https://jsonlogic.com/), which I basically use as a (restricted) AST for a SQL query. The server is written in golang.

# the name

It's kind of fun ... "sql" is a subsequence of "squirrel" ... "squirrel" means "to store up for future use" ... & that's what we're doing with our bytes ... the domain name was available
