# Go Ashot

This project is just a POC to learn about [Go Rod](https://go-rod.github.io/#/).

I also wanted to learn about [Go HTML Template](https://golang.org/pkg/html/template/) and
combine it with [Gin](https://github.com/gin-gonic/gin).

All of this glue together with [Uber FX](https://github.com/uber-go/fx) DI framework.

To make it more fun and useful, I decided to use it to crawl [Rappi](https://rappi.com.uy) website to place some orders.

> Yeah, we are a bunch of lazy people at the office, and we love Shawarmas 🥙.

## How to use it

Build the project:

```bash
make build
```

This will output two binaries, one for the server and another one for the crawler.

Run it:

```bash
./bin/app
#or
./bin/crawler
```

It will start a web server on port 5000, so you can access it on http://localhost:5000.

## How it works

The project is divided in two parts:

- The web server, which is just a simple web server that serves some Go HTML templates.
  With this web server, you can CRUD orders

- The crawler, which read the orders that were created on the web server, and place them on Rappi.
  Rappi has a 2FA, so it will take to the login page, and it will wait for you to complete the login process.
  Once you are logged in, it will place the orders.
