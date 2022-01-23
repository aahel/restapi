# restapi

Rest api without using any external web framework or router
## Run

`go run main.go`

## Usage

### Fetch data from mongodb

`POST` http://ec2-54-211-88-48.compute-1.amazonaws.com/v1/records

`Ex request:`
```
{
    "startDate": "2016-01-26",
    "endDate": "2016-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```
`Ex response:`

```
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "key": "bxoQiSKL",
      "totalCount": 2991,
      "createdAt": "2016-01-29T01:59:53.494Z"
    },
    {
      "key": "NOdGNUDn",
      "totalCount": 2813,
      "createdAt": "2016-01-28T07:10:33.558Z"
    }
  ]
}
```


### In memory endpoints

1. `POST` http://ec2-54-211-88-48.compute-1.amazonaws.com/v1/in-memory


Ex request and response:

```
{
    "key": "active-tabs",
    "value": "getir"
}
```

2. `GET` http://ec2-54-211-88-48.compute-1.amazonaws.com/v1/in-memory?key=(key)

Ex: `GET` http://ec2-54-211-88-48.compute-1.amazonaws.com/v1/in-memory?key=active-tabs

Ex response

```
{
    "key": "active-tabs",
    "value": "getir"
}
```

`Api Doc`

`GET` http://ec2-54-211-88-48.compute-1.amazonaws.com/docs