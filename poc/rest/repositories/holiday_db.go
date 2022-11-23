package repositories

import (
	"poc/model"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type holidayRepoDB struct {
	db  *gorm.DB
	rdb *redis.Client
}

type Tabler interface {
	TableName() string
}

func (Holiday) TableName() string {
	return "glo_holiday"
}

func NewHolidayRepoDB(db *gorm.DB, rdb *redis.Client) HolidayRepository {
	return holidayRepoDB{
		db:  db,
		rdb: rdb}
}

// func NewHolidayRepoRedis(rdb *redis.Client) HolidayRepository {
// 	return holidayRepoDB{rdb: rdb}
// }

func (r holidayRepoDB) GetHoliday() ([]Holiday, error) {
	holidays := []Holiday{}
	// result := r.db.Raw("SELECT * FROM glo_holiday").Scan(&holiday)
	result := r.db.Find(&holidays)
	// fmt.Println(result.Statement)
	if result.Error != nil {
		return nil, result.Error
	}
	return holidays, nil
}

func (r holidayRepoDB) AddHoliday(input model.HolidayRequest) error {
	holiday := Holiday{}
	result := r.db.FirstOrCreate(&holiday, input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r holidayRepoDB) GetRedis(key string) (string, error) {
	res, err := r.rdb.Get(key).Result()
	if err != nil {
		return "cannot get value from redis", err
	}
	return res, nil
}
