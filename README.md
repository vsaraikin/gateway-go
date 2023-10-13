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
4. Pass config with `API_KEY` and `SECRET_KEY` to `./config/.env` file.

## What's next?

1. Logging
2. Add `context.Context`
3. Errors handling and custom types
4. Add validation while sending request that there is no typo.


## References:

1. [How did I improve latency by 700% using sync.Pool](https://www.akshaydeo.com/blog/2017/12/23/How-did-I-improve-latency-by-700-percent-using-syncPool/)
2. [Detailed performance analysis of a simple low-latency trading system](https://sissoftwarefactory.com/blog/detailed-performance-analysis-of-a-simple-low-latency-trading-system/)