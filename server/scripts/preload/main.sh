#!/bin/bash
set -e

which curl
which jq

# oauth
OUT=$(curl \
--location --request PUT 'localhost:9922/api/oauth/providers/strava/config' \
--header 'Content-Type: application/json' \
--silent \
--data-raw '{
    "name": "strava",
    "client_id": "'$STRAVA_CLIENT_ID'",
    "client_secret": "'$STRAVA_CLIENT_SECRET'",
    "auth_url": "https://www.strava.com/api/v3/oauth/authorize",
    "token_url": "https://www.strava.com/api/v3/oauth/token",
    "redirect_url": "http://localhost:9921/oauth/providers/strava/token",
    "scopes": [
        "activity:readall"
    ],
    "auth_url_params": {
        "approval_prompt": "force"
    }
}')

echo $OUT | jq '.client_secret = "********"'

OUT=$(curl \
--location --request PUT 'localhost:9922/api/oauth/providers/spotify/config' \
--header 'Content-Type: application/json' \
--silent \
--data-raw '{
    "name": "spotify",
    "client_id": "'$SPOTIFY_CLIENT_ID'",
    "client_secret": "'$SPOTIFY_CLIENT_SECRET'",
    "auth_url": "https://accounts.spotify.com/authorize",
    "token_url": "https://accounts.spotify.com/api/token",
    "redirect_url": "http://localhost:9921/oauth/providers/spotify/token",
    "scopes": [
        "user-read-recently-played",
	    "user-top-read",
	    "playlist-read-private",
	    "playlist-read-collaborative",
	    "user-follow-read",
	    "user-library-read"
    ]
}')

echo $OUT | jq '.client_secret = "********"'

curl \
--location --request PUT 'localhost:9922/api/kvs/hello' \
--header 'Content-Type: application/json' \
--silent \
--data-raw '{
    "key": "hello",
    "value": "world"
}' | jq '.'

curl --location --request PUT 'localhost:9922/api/documents/hello' \
--header 'Content-Type: application/json' \
--silent \
--data-raw '{
    "id": "hello",
    "body": {
        "foo": "bar"
    },
    "header": {
        "fizz": "buzz"
    }
}' | jq '.'
