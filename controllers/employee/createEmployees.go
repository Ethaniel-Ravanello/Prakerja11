package employee

import (
	"keterampilan/config"
	"keterampilan/models/base"
	"keterampilan/models/employee"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateEmployee(c echo.Context) error {
	var employee employee.Employee
	if err := c.Bind(&employee); err != nil {

		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Error:   true,
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	if err := config.DB.Create(&employee).Error; err != nil {

		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Error:   true,
			Code:    http.StatusInternalServerError,
			Message: "Failed to create employee",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Error:   false,
		Code:    http.StatusCreated,
		Message: "Employee created successfully",
		Data:    employee,
	})
}
