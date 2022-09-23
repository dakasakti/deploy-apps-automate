package main

import (
	"fmt"

	"github.com/dakasakti/deploy-apps-hexagonal/config"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/factory"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/http"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/routes"
)

func main() {
	c := config.GetConfig()
	f := factory.NewFactory(c)

	e := routes.New()
	middlewares.LoggerMiddleware(e)

	http.NewHttp(e, f)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
