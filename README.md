# URL Shortener

Installation.
make sure you have installed all your mod dependencies first.

Setup Enviroment variables.

```
export $(cat .env)
```

The run server.

```
go run main.go
```

to try msgpack please go to tool and then run:

```
go run msgpack.go
```


POST http://localhost:9000/

## json body
```
{
   "url":"https://google.com"
}
```
## response
```
{
    "code": "eSshzhbGg",
    "url": "http://google.com",
    "created_at": 1608844847
}
```

GET http://localhost:9000/eSshzhbGg
## response
```
Will render http://google.com
```

