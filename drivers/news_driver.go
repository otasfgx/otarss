package drivers

import (
	"context"
	"fmt"
	"os"
	"rss/adapters/controllers"

	"github.com/labstack/echo"
)

type News interface {
	Run(ctx context.Context)
}

type NewsDriver struct {
	echo       *echo.Echo
	controller controllers.News
}

func NewNewsDriver(echo *echo.Echo, controller controllers.News) News {
	return &NewsDriver{
		echo:       echo,
		controller: controller,
	}
}

func (driver *NewsDriver) Run(ctx context.Context) {
	driver.echo.POST("/", driver.controller.GetNews(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
