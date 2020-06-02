# SLIC
SLIC is a simple location indexing service built with golang and persisted with redis for speed.
It's ingests lng/lats and returns a geoIndex. 
The geospatial indexing library being used here was built by uber [h3](https://github.com/uber/h3-go)
Feel free to use it.

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
        "driver_id": "1",
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
