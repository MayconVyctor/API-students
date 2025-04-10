package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type StudentResponse struct {
	ID        int       `json:"id" example:"1"`
	CreatedAt time.Time `json:"createdAt" example:"2025-03-31T12:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-03-31T12:30:00Z"`
	DeletedAt time.Time `json:"deletedAt" example:"0001-01-01T00:00:00Z"` // zero time
	Name      string    `json:"name" example:""`
	CPF       int       `json:"cpf" example:""`
	Email     string    `json:"email" example:""`
	Age       int       `json:"age" example:""`
	Active    bool      `json:"active" example:""`
}

func NewResponse(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students {
		studentResponse := StudentResponse{
			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Name:      student.Name,
			Email:     student.Email,
			CPF:       student.CPF,
			Age:       student.Age,
			Active:    student.Active,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}
	return studentsResponse
}
