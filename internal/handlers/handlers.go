package handlers

import (
	"net/http"
	"net/netip"

	"github.com/labstack/echo/v4"

	db "github.com/rem-aster/goth-template/internal/database"
	v "github.com/rem-aster/goth-template/internal/views"
	vc "github.com/rem-aster/goth-template/internal/views/components"
)

func (h *Handler) Index(e echo.Context) error {
	ip, err := netip.ParseAddr(e.RealIP())
	if err != nil {
		return err
	}
	user, err := h.db.GetUserByIP(e.Request().Context(), &ip)
	var username string
	if err != nil {
		user, err = h.db.NewUser(e.Request().Context(), &ip)
		if err != nil {
			return err
		}
	}
	username = user.Username
	return h.render(e, http.StatusOK, v.Index(username))
}

func (h *Handler) ChangeUsername(e echo.Context) error {
	username := e.FormValue("username")
	ip, err := netip.ParseAddr(e.RealIP())
	if err != nil {
		return err
	}
	_, err = h.db.ChangeUsername(e.Request().Context(), db.ChangeUsernameParams{
		Username: username,
		Ip:       &ip,
	})
	if err != nil {
		return err
	}
	return h.render(e, http.StatusOK, vc.AlertSuccess("Username changed successfully!"))
}
