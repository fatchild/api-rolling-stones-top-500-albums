# Rolling Stones Top 500 Albums API

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
