package main

import (
	"IrisProduct/backend/web/controllers"
	"IrisProduct/common"
	"IrisProduct/repository"
	"IrisProduct/services"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

func main() {
	app := iris.New()
	// 设置log
	app.Logger().SetLevel("debug")
	// 设置iris模板
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	// 设置模板目标
	// 使用HandleDir代替StaticWeb
	app.HandleDir("/assets", "./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Println(err.Error())
	}
	//??
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//5.注册控制器
	productRepository := repository.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
