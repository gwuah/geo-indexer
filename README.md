# Athena

Athena is our highly performant in-memory driver location indexing service. 
It's built specifically to receive geo-coordinates from drivers and index them for faster lookups during the dispatching phase of our lifecycle.
It's usage may change in the near future

# How to run 
First Way
- `go run server.go`

Second Way
- `go build server.go` and then you run `./server`

# Usage

To index driver location data,

`POST /index-driver-location`

```
{
    "driver_id": "1",
    "lat": "5.68662590662494",
    "lng": "-0.24954311060899848"
}
```

should return 
When the driver hasn't changed their zone.

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
        "last_driver_index": "614792924527329279",
        "latest_driver_index": 614792923793326079
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
