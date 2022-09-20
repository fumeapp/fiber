# Fume adapter for Fiber

[![Go Reference][1]][2]
[![GoCard][3]][4]

[1]: https://pkg.go.dev/badge/github.com/fumeapp/fiber.svg
[2]: https://pkg.go.dev/github.com/fumeapp/fiber
[3]: https://goreportcard.com/badge/github.com/fumeapp/fiber
[4]: https://goreportcard.com/report/github.com/fumeapp/fiber

This is a simple adapter for [Fiber](https://gofiber.io/) that allows you to deploy your application using Fume.

```go
package main

import (
	fume "github.com/fumeapp/fiber"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()
	app.GET("/", func(c *fiber.Ctx) { c.Status(200).JSON(&fiber.Map{"message": "Hello World"}) })
	fume.Start(app, fume.Options{})
}
```

## Options

| Option | Description | Default |
| ------ | ----------- | ------- |
| `Port` | Port to listen on | `8000` |
| `Host` | Host to listen on | `localhost` |
