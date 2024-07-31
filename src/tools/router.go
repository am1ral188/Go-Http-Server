package tools

import (
	"fmt"
	"net/http"
)

type pro struct {
	Method        string
	Controller    Controller
	Act           string
	HandleSubPath bool
}
type action map[string]pro
type callBack func(*Router)

type Router struct {
	middleWare    func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
	hasMiddleWare bool
	Path          string
	Children      []*Router
	Actions       action
}

func (r *Router) Group(path string, c ...callBack) *Router {
	r.Actions = make(action, 0)
	r.Path = path
	obj := &Router{}
	obj.Actions = make(action, 0)
	r.Children = append(r.Children, obj)
	c[0](obj)
	return obj
}
func (r *Router) UseMiddleWare(f func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	r.middleWare = f
	r.hasMiddleWare = true
}
func (r *Router) GET(path string, cont Controller, act string, handleSubPath bool) {
	r.Actions[path] = pro{"GET", cont, act, handleSubPath}
}
func (r *Router) POST(path string, cont Controller, act string, handleSubPath bool) {
	r.Actions[path] = pro{"POST", cont, act, handleSubPath}
}
func (r *Router) init(obj *Router, parentPath string, writer http.ResponseWriter, request *http.Request, hasM bool) bool {
	notFount := true
	for s, f := range obj.Actions {
		if hasM {
			if HandleWithMiddleWare(parentPath+s, f.Controller, f.Act, f.Method, writer, request, r.middleWare, f.HandleSubPath) {
				notFount = false
				return notFount
			}
		} else {
			if Handle(parentPath+s, f.Controller, f.Act, f.Method, writer, request, f.HandleSubPath) {
				notFount = false
				return notFount
			}
		}

	}

	p := parentPath + obj.Path
	for _, child := range obj.Children {
		return r.init(child, p, writer, request, child.hasMiddleWare)
	}
	return notFount
}
func (r *Router) Init() {
	Mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		res := r.init(r, "", writer, request, r.hasMiddleWare)
		if res {
			view, err := View("notFound")
			if err != nil {
				fmt.Println(err)
				return
			}
			err2 := view.Show(writer)
			if err2 != nil {
				fmt.Println()
				return
			}
		}
	})

}
