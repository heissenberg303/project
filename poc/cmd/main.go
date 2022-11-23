package main

import (
	"fmt"
	"poc/handlers"
	"poc/repositories"
	"poc/services"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("POC of Hexagonal Architecture start...")
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dsn := "root:P@ssw0rd@tcp(localhost:3306)/ktb_glo?parseTime=true"
	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	holidayRepo := repositories.NewHolidayRepoDB(db, redisDB)
	holidayService := services.NewHolidayService(holidayRepo)
	holidayHandler := handlers.NewHolidayHandler(holidayService)
	e := echo.New()
	e.GET("/", holidayHandler.GetAllHoliday)
	e.Start(":80")

	// keyCurrentPeriod := "CURRENTPERIOD"
	// currentPeriod, err := rdb.Get(keyCurrentPeriod).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(currentPeriod)
	// redisRepo := repositories.NewHolidayRepoRedis(rdb)
	// res, err := redisRepo.GetRedis(keyCurrentPeriod)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)

	// holidayHandler := handlers.NewHolidayHandler(holidayService)
	// e := echo.New()
	// e.GET("/", holidayHandler.GetAllHoliday)
	// e.POST("/add", holidayHandler.AddHoliday)

	// r := router.InitRoute()
	// e.Start(":80")

	// username:password@protocol(address)/dbname?param=value
	// dsn := "root:P@ssw0rd@tcp(localhost:3306)/ktb_glo?parseTime=true"
	// dial := mysql.Open(dsn)

	// db, err := gorm.Open(dial)
	// if err != nil {
	// 	panic(err)
	// }

	// holidayRepo := repositories.NewHolidayRepoDB(db)
	// // result, err := holidayRepo.GetHoliday()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result)

	// mockHoliday := repositories.Holiday{
	// 	Id:     49,
	// 	Year:   "2023",
	// 	Date:   "2023-09-04",
	// 	Detail: "Cat Day",
	// }
	// err = holidayRepo.AddHoliday(mockHoliday)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("SUCCEED")
	// holidayService := services.NewHolidayService(holidayRepo)
	// holidayHandler := handlers.NewHolidayHandler(holidayService)
	// e := echo.New()
	// e.GET("/", holidayHandler.GetAllHoliday)
	// e.Start(":80")
	// fmt.Println("routing at port 80")
}
