package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MayconVyctor/API-students/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// getStudents godoc
//
// @Summary      Get a list of students
// @Description  Retrieve students details
// @Tags         students
// @Accept       json
// @Produce      json
// Param         register path int false "Registration"
// @Success      200 {object} schemas.StudentResponse
// @Failure      404
// @Router       /students [get]
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	active := c.QueryParam("active")

	if active != "" {
		act, err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to parse boolean")
			return c.String(http.StatusInternalServerError, "Failed to parse boolean")

		}
		students, err = api.DB.GetFilteredStudents(act)
	}

	listOfStudents := map[string][]schemas.StudentResponse{"students": schemas.NewResponse(students)}

	return c.JSON(http.StatusOK, listOfStudents)
}

// createStudent godoc
//
// @Summary      Create student
// @Description  Create and store a new student in the database
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      201 {object} schemas.StudentResponse
// @Failure      400
// @Router       /students [post]
func (api *API) CreateStudent(c echo.Context) error {
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validating struct")
		return c.String(http.StatusBadRequest, "Error validating student")
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		Email:  studentReq.Email,
		CPF:    studentReq.CPF,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create Student")
	}
	return c.JSON(http.StatusOK, student)
}

// getStudent godoc
//
// @Summary      Get student by ID
// @Description  Retrieve student details using the student's ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200 {object} schemas.StudentResponse
// @Failure      404
// @Failure      500
// @Router       /students/{id} [get]
func (api *API) GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}

	student, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	return c.JSON(http.StatusOK, student)
}

// updateStudent godoc
//
// @Summary      Update an existing student
// @Description  Update a student's information using their ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200 {object} schemas.StudentResponse
// @Failure      404
// @Failure      500
// @Router       /students/{id} [put]
func (api *API) UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student := UpdateStudentInfo(receivedStudent, updatingStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}

	return c.JSON(http.StatusOK, student)
}

func UpdateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}

	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}

	if receivedStudent.CPF > 0 {
		student.CPF = receivedStudent.CPF
	}

	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}

	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}
	return student
}

// deleteStudent godoc
//
// @Summary      Delete a student
// @Description  Remove a student from the database using their ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200 {object} schemas.StudentResponse
// @Failure      404
// @Failure      500
// @Router       /students/{id} [delete]
func (api *API) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid student ID")
	}

	student, err := api.DB.GetStudent(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Student not found")
		}
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, student)
}
