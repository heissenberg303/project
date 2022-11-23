package handlers

import (
	"net/http"
	"poc/model"
	"poc/services"

	"github.com/labstack/echo/v4"
)

type holidayHandler struct {
	holidaySrv services.HolidayService
}

func NewHolidayHandler(holidaySrv services.HolidayService) holidayHandler {
	return holidayHandler{holidaySrv: holidaySrv}
}

func (h holidayHandler) GetAllHoliday(c echo.Context) error {
	holidays, err := h.holidaySrv.GetAllHoliday()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, "")
	}

	// response, err := json.Marshal(holidays)
	// if err != nil {
	// 	return c.JSONPretty(http.StatusInternalServerError, err, "cannot marshal json")
	// }
	return c.JSONPretty(http.StatusOK, holidays, "")
}

func (h holidayHandler) AddHoliday(c echo.Context) error {

	req := new(model.HolidayRequest)
	err := c.Bind(req)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, "can not bind request")
	}
	err = h.holidaySrv.AddHoliday(*req)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, "call AddHoliday error")
	}

	return c.JSONPretty(http.StatusOK, "sucess", "")
}
