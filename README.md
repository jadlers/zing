# Zzng [![Go Report Card](https://goreportcard.com/badge/github.com/jadlers/zzng)](https://goreportcard.com/report/github.com/jadlers/zzng)

A simple CLI for getting the lyrics for songs those time you just have to sing
along to the music blasting around you. Sing out more, sing out loud 🔊

## Build/Installation

You will need API keys for both Genius ([create
here](https://genius.com/api-clients/new)) and apiseeds ([create
here](https://apiseeds.com)) in order to compile.

To make it easier to hack on the project create a `Makefile` with the following
content and add your keys:

```makefile
GENIUS_CLIENT_ACCESS_TOKEN=
APISEEDS_API_KEY=

ldflags="-X github.com/jadlers/zzng/genius.clientAccessToken=$(GENIUS_CLIENT_ACCESS_TOKEN)\
         -X github.com/jadlers/zzng/apiseeds.clientApiKey=$(APISEEDS_API_KEY)"

build:
	go build -ldflags=${ldflags}

install:
	go install -ldflags=${ldflags}
```

Now all you need to do is run `make build` or `make install`.
