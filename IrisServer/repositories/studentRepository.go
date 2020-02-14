package repositories

import (
	"IrisServer/models"
	_ "fmt"
)

type StudentRepository interface {
	GetName() string
}

type StudentManager struct {
}

func NewStudentManager() StudentRepository {
	return &StudentManager{}
}

func (m *StudentManager) GetName() string {
	s := &models.Student{
		Name: "Alex",
		Age:  10,
	}
	return s.Name
}
