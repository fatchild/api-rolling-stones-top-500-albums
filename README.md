# Rolling Stones Top 500 Albums API

Live instance version: 0.0.1-alpha - https://api-rolling-stones-top-500-albums.fly.dev/

## Setup

1. Clone the repo

`git clone https://github.com/fatchild/api-rolling-stones-top-500-albums.git`

2. Install dependencies

`go mod tidy`

3. Test the installation

`go test ./...`

4. Add a `.env` file for test environment

```
URL=localhost
PORT='8080'
LOG_TO_FILE=TRUE
RELEASE_MODE=TEST
```

5. Run the project

`go run .`

## API

> Get all albums on the list (2024)
> The API call returns the list of albums

`curl http://localhost:8080/getAlbumList`

```json
[
{
  "position": 1,
  "album": "What's Going On",
  "artist": "Marvin Gaye",
},
{
  "position": 500,
  "album": "Funeral",
  "artist": "Arcade Fire",
}
]
```

> Get a specific album in the list (2024) from it's position
> The API call returns a single album

`curl http://localhost:8080/getAlbumAtPosition/7`

```json
{
  "position": 7,
  "artist": "Fleetwood Mac",
  "album": "Rumours"
}
```
