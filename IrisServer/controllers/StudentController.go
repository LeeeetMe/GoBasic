package controllers

import (
	"IrisServer/repositories"
	services2 "IrisServer/services"
	"github.com/kataras/iris/mvc"
)

type StudentController struct {

}

func (s* StudentController) Get() mvc.View {
	repo := repositories.NewStudentManager()
	ser := services2.NewStudentServiceManager(repo)
	name := ser.GetStudentName()
	return mvc.View{
		Data:name,
		Name:"studentIndex.html",
	}
}