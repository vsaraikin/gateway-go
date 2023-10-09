# Gateway for Binance exchange

TL;DR Gateway to Binance that handles market data.

Now it handles only `@depth`, but as I have made handling websocket connection, it is easy to scale it to other endpoints as well as to the other exchanges.

**Run with Docker:** 

```shell
make app
```

**Run locally:** 

```shell
go run main.go
```

## Features

1. Graceful shutdown – all connections are ended smoothly.
2. Ready to scale to other endpoints.
3. Runs each subscription in a different goroutine (thread).
4. Pass config with `API_KEY` and `SECRET_KEY` to `./config/.env file`.
