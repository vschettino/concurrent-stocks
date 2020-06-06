# Concurrent stock price fetcher

This repository has some simple examples on how to wait for [goroutines](https://golangdocs.com/goroutines-in-golang).

More about this examples in this [blog post](https://vschettino.me/posts/waiting-go-routines-fetch-stock-prices).

## Dependencies
This is a wrap up of the awesome [go-quote](https://github.com/markcheno/go-quote).

## Install and run

1. Clone this repository
1. `go run main.go`

I used the most simple benchmark of them all: Linux's [time function](http://manpages.ubuntu.com/manpages/bionic/en/man7/time.7.html). To do the same you can call the program like:

```shell script
time go run main.go
```
