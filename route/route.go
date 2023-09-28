package routes

import (
	"keterampilan/controllers/employee"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/employees", employee.GetEmployeeList)
	e.GET("/employees/:id", employee.GetEmployeeListById)
	e.POST("/employees", employee.CreateEmployee)
}
