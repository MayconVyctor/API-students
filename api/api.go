package api

import (
	"github.com/MayconVyctor/API-students/db"

	_ "github.com/MayconVyctor/API-students/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) ConfigureRoutes() {

	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.CreateStudent)
	api.Echo.GET("/students/:id", api.GetStudent)
	api.Echo.PUT("/students/:id", api.UpdateStudent)
	api.Echo.DELETE("/students/:id", api.DeleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}
