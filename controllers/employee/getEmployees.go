package employee

import (
	"keterampilan/config"
	"keterampilan/models/base"
	"keterampilan/models/employee"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployeeList(c echo.Context) error {
	var employeeList []employee.Employee

	config.DB.Find(&employeeList)
	return c.JSON(http.StatusOK, base.BaseResponse{
		Error:   false,
		Code:    200,
		Message: "Success Get Data",
		Data:    employeeList,
	})
}
