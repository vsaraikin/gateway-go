# Gateway exchange

TL;DR Gateway to Binance that handles market data. 

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
