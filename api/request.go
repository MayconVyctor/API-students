package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name" example:""`
	CPF    int    `json:"cpf" example:""`
	Email  string `json:"email" example:""`
	Age    int    `json:"age" example:""`
	Active *bool  `json:"active" example:""`
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("param '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("cpf", "int")
	}
	if s.Age == 0 {
		return errParamRequired("age", "int")
	}
	if s.Active == nil {
		return errParamRequired("active", "booll")
	}
	return nil
}
