# deezer-bao-playlist-generator

A Golang based project that can add the top 3 songs to [Deezer](https://deezer.com/) playlist for each band that is playing at this years [Bloodstock Open Air festival](https://www.bloodstock.uk.com/).

To use this app you will need to get a Deezer access_token and the ID of an existing Deezer playlist.

## Build & Run

Build:

`go build main.go`

Run:

`DEEZER_ACCESS_TOKEN=<TOKEN> DEEZER_PLAYLIST_ID=<TOKEN> ./deezer-boa-playlist-generator`

## Get an Deezer Access token

See Deezer's [OAuth developer](https://developers.deezer.com/api/oauth) guide 
