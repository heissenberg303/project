package services

import "poc/model"

type HolidayService interface {
	GetAllHoliday() ([]model.HolidayResponse, error)
	AddHoliday(model.HolidayRequest) error
	// GetRedis(key string) (string, error)
}
