package Controller

import (
	"awesomeProject/src/tools"
	"fmt"
	"log"
	"net/http"
)

type MainController struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (r MainController) Index() {
	if r.Request.URL.String() != "/" {
		view, err := tools.View("notfound")
		if err != nil {
			log.Fatal(err)
		}
		err2 := view.Show(r.Response)
		if err2 != nil {
			log.Fatal(err2)
		}
	} else {
		view, err := tools.View("index")
		if err != nil {
			fmt.Println(err)
		}
		view.Show(r.Response)
	}

}

func (r *MainController) Set(w http.ResponseWriter, req *http.Request) {
	r.Response = w
	r.Request = req

}
func NewMainController() *MainController {
	return &MainController{}
}
