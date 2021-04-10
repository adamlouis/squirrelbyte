# server

`server` is the squirrelbyte server

it's a json blob server

it supports a query syntax similar to [jsonlogic](https://jsonlogic.com/)

it uses sqlite as a data store

the `/documents:search` endpoint DOES NOT use prepared queries, so it's recommended `/documents:search` is used only in safe contexts (read-only, local)

feels verbose ... too much indirection ... but the show must go on

it serves the static web bundle

perhaps more to come ... getting it out there

