package router

import (
	Controller "awesomeProject/src/controller"
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
	r.Init()
}
