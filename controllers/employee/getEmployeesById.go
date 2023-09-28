package employee

import (
	"keterampilan/config"
	"keterampilan/models/base"
	"keterampilan/models/employee"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployeeListById(c echo.Context) error {
	employeeID := c.Param("id")

	var employee employee.Employee

	if err := config.DB.Where("id = ?", employeeID).First(&employee).Error; err != nil {

		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Error:   true,
			Code:    http.StatusNotFound,
			Message: "Employee not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Error:   false,
		Code:    http.StatusOK,
		Message: "Success Get Employee by ID",
		Data:    employee,
	})
}
