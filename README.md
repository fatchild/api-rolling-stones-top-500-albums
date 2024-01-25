# Rolling Stones Top 500 Albums API

## Useful Articles

- [RHEL - What is a REST API?](https://www.redhat.com/en/topics/api/what-is-a-rest-api)
- [Creating Golang API with Gin Gonic and GCP Firebase (Part 1) - Go Tutorial](https://www.youtube.com/watch?v=UlZ_EGWvN7w&ab_channel=MarcioMarinho)
- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
- [What are the types of APIs and their differences?](https://www.techtarget.com/searchapparchitecture/tip/What-are-the-types-of-APIs-and-their-differences)

## Goal

Get information about the various versions of the RS top 500 album lists.

1. Making a request such as `/getAlbumList?all=true&year=2024&listBy=position` would get you the full 2024 list of albums.

2. Client receives the following;

```js
{
    [
        {
            album: "What's Going On",
            artist: "Marvin Gaye",
        }, 
        {
            album: "Pet Sounds",
            artist: "The Beach Boys",
        }, 
        ...
        {
            album: "Funeral",
            artist: "Arcade Fire",
        }, 
        
    ]
}
```

# Notes

## API TYPE

A public API which can be accesses through the public internet.

## API Architecture & Protocol

RESTful API is a Representational State Transfer API. Transferring the state of the program in a standardized format.

- Stores no data between requests (stateless)
- Mainly used for exchanging data
- RPC (Remote Procedural Call) This is used for action invocation
- SOAP (Simple Object Access Protocol) a tightly controlled standard

