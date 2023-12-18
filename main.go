package main

import (
	"context"
	"net/http"

	"github.com/leifarriens/go-templ/app/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Index("You"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil

}
