# SLIC
SLIC is a simple location indexing service built with golang and persisted with redis.
It's ingests lng/lats and returns a geoIndex. 
The geospatial indexing library being used here was built by uber [h3](https://github.com/uber/h3-go)
Feel free to use it.

# What it does
It takes several coordinates and places them under an index, such that 1 index can point to numerous user ids.

# What can you do with this (build upon this)
- You can build a realtime supply service for your ride sharing application
- You can build a realtime demand service for your ride sharing application
- Build a price surging system based on the number of id's in a particular geographical area (geo-index)

# How to run 
First Way
- `go run server.go`

Second Way
- `go build server.go` and then you run `./server`

# Usage

To index location data,

`POST /index-location`

```
{
    "id": "1",
    "lat": "5.68662590662494",
    "lng": "-0.24954311060899848"
}
```

should return 
When the user hasn't changed their zone.

```
{
    "message": "Success"
}
```

When the driver has entered a new zone
```
{
    "data": {
        "id": "1",
        "last_index": "614792924527329279",
        "latest_index": 614792923793326079
    },
    "message": "Success"
}
```
When there's an error
```
{
    "message": "Error"
}
```
