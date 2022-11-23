package router

import (
	"poc/handlers"
	"poc/repositories"
	"poc/services"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitRoute() *echo.Echo {

	dsn := "root:P@ssw0rd@tcp(localhost:3306)/ktb_glo?parseTime=true"
	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	holidayRepo := repositories.NewHolidayRepoDB(db)
	holidayService := services.NewHolidayService(holidayRepo)
	holidayHandler := handlers.NewHolidayHandler(holidayService)
	e := echo.New()
	e.GET("/", holidayHandler.GetAllHoliday)
	e.POST("/add", holidayHandler.AddHoliday)
	e.Start(":80")

	return e
}
