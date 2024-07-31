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
		router.GET("/card-verify", c, "CardVerify", false)
		router.GET("/new-card", c, "CardInsert", false)
		router.GET("", c, "Index", false)
	})
	r.GET("/", Controller.NewMainController(), "Index", false)
	r.GET(`/{id:^[0-9]+$}`, Controller.NewMainController(), "ID", true)
	r.UseMiddleWare(middleWares.ExampleMiddleWare)
	r.Init()
}
