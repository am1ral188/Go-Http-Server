package Controller

import (
	"awesomeProject/src/tools"
	"fmt"
	"io"
	"net/http"
)

type MainController struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (r MainController) Index(args ...interface{}) {
	view, err := tools.View("index")
	if err != nil {
		fmt.Println(err)
	}
	view.Show(r.Response)

}
func (r MainController) ID(id string) {
	io.WriteString(r.Response, id)
}

func (r *MainController) Set(w http.ResponseWriter, req *http.Request) {
	r.Response = w
	r.Request = req

}
func NewMainController() *MainController {
	return &MainController{}
}
