package controllers

import (
	"IrisProduct/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.ProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray, _:= p.ProductService.SelectAll()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"Products": productArray,
		},
	}
}
