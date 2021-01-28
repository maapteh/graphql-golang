# Go implementation GraphQL
> simple Golang setup with modules utilizing GraphQL

## PRE-REQUISITES
- have golang installed locally for development


## DEVELOPMENT
- run `PORT=8080 GIN_MODE=debug go run server.go` or TODO HOT RELOAD
- generate `gqlgen generate`
- playground is available in `http://localhost:8080/`, you can turn off polling with settings and set `"schema.polling.enable": false,`

### LIVE-RELOAD
- install from `https://github.com/cosmtrek/air`

curl example:
```
curl \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{ "query": "{ crocodiles { age } }" }' \
  http://localhost:8080/graphql
```


load test run with hey:
```
hey -n 200 -c 80 -cpus 1 -z 10s -m POST --disable-keepalive -H "Content-Type: application/json" -d '{ "query": "{ crocodiles { age } }" }' http://localhost:8080/graphql
```

## EXAMPLES
simple example query you can run in our playground:
```
{
  bert: crocodile(id: 1) {
    id
    name
    age
  }
  tom: crocodile(id: 2) {
    name
    sex
  }
  bar:crocodiles {
    name
  }
  gone:crocodile(id: 7) {
    id
  }
} 
```


## HEALTH CHECK
- simple endpoint: 'http://localhost:8080/.well-known/go/server-health'

```
curl -X GET --head -i http://localhost:8080/.well-known/server-health
HTTP/1.1 200 OK
Date: Sat, 10 Oct 2020 17:35:43 GMT
Content-Length: 0
```

## CREDITS

- https://github.com/bv1p Goran
- https://github.com/maapteh Maarten
