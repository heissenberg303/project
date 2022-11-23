package repositories

import "poc/model"

type HolidayRepository interface {
	GetHoliday() ([]Holiday, error)
	AddHoliday(model.HolidayRequest) error
	GetRedis(key string) (string, error)
}

type Holiday struct {
	Id     int    `gorm:"id"`
	Year   string `gorm:"year"`
	Date   string `gorm:"date"`
	Detail string `gorm:"detail"`
}
