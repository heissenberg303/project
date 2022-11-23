package services

import (
	"poc/model"
	"poc/repositories"
)

type holidayService struct {
	holidayRepo repositories.HolidayRepository
}

func NewHolidayService(holidayRepo repositories.HolidayRepository) HolidayService {
	return holidayService{holidayRepo: holidayRepo}
}

// func (r holidayService) GetRedis(key string) (string, error) {

// 	return "", error
// }

func (r holidayService) GetAllHoliday() ([]model.HolidayResponse, error) {
	response := []model.HolidayResponse{}
	dataRedis, err := r.holidayRepo.GetRedis("CURRENTPERIOD")
	if err != nil {
		result, errDB := r.holidayRepo.GetHoliday()
		if errDB != nil {
			return nil, errDB
		}
		for _, res := range result {
			data := model.HolidayResponse{
				Date:   res.Date,
				Detail: res.Detail,
			}
			response = append(response, data)
		}
		return response, nil
	}
	// res := []model.HolidayResponse{}
	res := make([]model.HolidayResponse, 3)

	for i := 0; i < 3; i++ {
		res[i] = model.HolidayResponse{
			Date:   string(dataRedis),
			Detail: "get date from redis",
		}
	}
	return res, nil

	// if err != nil {
	// 	return nil, err
	// }

}

func (r holidayService) AddHoliday(req model.HolidayRequest) error {

	// r.holidayRepo.AddHoliday()
	err := r.holidayRepo.AddHoliday(req)
	if err != nil {
		return err
	}
	return nil
}
