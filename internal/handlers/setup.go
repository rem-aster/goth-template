package handlers

import (

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	db "github.com/rem-aster/goth-template/internal/database"
)

type Handler struct {
	srv *echo.Echo
	db  *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{
		srv: echo.New(),
		db:  db,
	}
}

func (h *Handler) SetupRoutes() {
	h.srv.Use(middleware.Logger())
	h.srv.Use(middleware.Recover())

	h.srv.Static("/", "web/static")

	h.srv.GET("/", h.Index)
	h.srv.POST("/change-username", h.ChangeUsername)
}

func (h *Handler) StartServer() error {
	err := h.srv.Start(":8081")
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) render(e echo.Context, code int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := component.Render(e.Request().Context(), buf); err != nil {
		return err
	}
	return e.HTML(code, buf.String())
}
