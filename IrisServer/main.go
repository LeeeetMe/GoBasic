package main

import (
	"IrisServer/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main()  {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views",".html").Reload(true))
	app.Logger().SetLevel("debug")
	// 添加controllers
	mvc.New(app.Party("/hello")).Handle(new(controllers.StudentController))
	_ = app.Run(iris.Addr("localhost:8080"))
}