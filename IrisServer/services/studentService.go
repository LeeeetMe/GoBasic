package services

import (
	"IrisServer/repositories"
)

type StudentService interface {
	GetStudentName() string
}

type StudentServiceManager struct {
	repo repositories.StudentRepository
}

func NewStudentServiceManager(repo repositories.StudentRepository) StudentService {
	s := &StudentServiceManager{repo: repo}
	return s
}

func (s *StudentServiceManager) GetStudentName() string {
	return "Student 获取的名字为" + s.repo.GetName()
}
