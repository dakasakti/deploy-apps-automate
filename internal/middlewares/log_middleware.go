package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, url=${host}${uri}, status=${status}, error=${error}, latency=${latency_human}\n",
	}))
}
