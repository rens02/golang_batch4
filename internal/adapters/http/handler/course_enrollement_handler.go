package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type CourseEnrollmentHandler struct {
	CourseEnrollmentUseCase usecase.CourseEnrollmentUseCase
}

func (handler CourseEnrollmentHandler) GetAllStudents() echo.HandlerFunc {
	return func(e echo.Context) error {
		var students []entity.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		students, err = handler.CourseEnrollmentUseCase.GetStudents(id) // Menggunakan parameter 'id' yang diterima dari URL
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		data := make(map[string]interface{})
		data["students"] = students
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status_code": http.StatusCreated,
			"message":     "The students has been successfully displayed",
			"data":        data,
		})
	}
}
func (handler CourseEnrollmentHandler) GetAllCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var courses []entity.Course
		sellerID, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		courses, err = handler.CourseEnrollmentUseCase.GetCourse(sellerID) // Menggunakan parameter 'id' yang diterima dari URL
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		data := make(map[string]interface{})
		data["courses"] = courses
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status_code": http.StatusCreated,
			"message":     "The course has been successfully displayed",
			"data":        data,
		})
	}
}
