package model

type HolidayRequest struct {
	Id     int    `json:"id"`
	Year   string `json:"year"`
	Date   string `json:"date"`
	Detail string `json:"detail"`
}

type HolidayResponse struct {
	Date   string `json:"date"`
	Detail string `json:"detail"`
}
