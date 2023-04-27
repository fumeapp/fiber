package fume

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

type Options struct {
	// optional - hostname to listen on (default: localhost)
	Host string
	// optional - port to run on (default: 8000)
	Port string
}

func Start(app *fiber.App, options Options) {

	defaults := &Options{
		Host: "localhost",
		Port: "8000",
	}

	if options.Port != "" {
		defaults.Port = options.Port
	}
	if options.Host != "" {
		defaults.Host = options.Host
	}

	if os.Getenv("_HANDLER") != "" {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		log.Fatal(app.Listen(defaults.Host + ":" + defaults.Port))
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, req)
}
