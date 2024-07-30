package router

import (
	Controller "awesomeProject/src/controller"
	"awesomeProject/src/router/middleWares"
	Tools "awesomeProject/src/tools"
)

func RouterInit() {
	r := Tools.Router{}
	r.Group("/api", func(router *Tools.Router) {
		c := Controller.NewApiController()
		router.GET("/card-verify", c, "CardVerify")
		router.GET("/new-card", c, "CardInsert")
		router.GET("", c, "Index")
	})
	r.GET("/", Controller.NewMainController(), "Index")
	r.GET(`/{id:^[0-9]+$}`, Controller.NewMainController(), "ID")
	r.UseMiddleWare(middleWares.ExampleMiddleWare)
	r.Init()
}
