package handler

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render is a wrapper around Echo's echo.Context.Render() with templ's templ.Component.Render()
func Render(ctx echo.Context, statusCode int, t ...templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := chain(t...).Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func chain(components ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		for _, c := range components {
			if err = c.Render(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
}
