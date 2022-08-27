package drivers

import (
	"context"
	"rss/adapters/controllers"

	"github.com/labstack/echo"
)

type News interface {
	Run(ctx context.Context, port string)
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

func (driver *NewsDriver) Run(ctx context.Context, port string) {
	driver.echo.POST("/", driver.controller.GetNews(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(port))
}
