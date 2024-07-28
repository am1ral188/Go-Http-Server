package tools

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Controller interface {
	Index()
	Set(http.ResponseWriter, *http.Request)
}

var Mux = http.NewServeMux()

func Handle(path string, cont Controller, act string, method string) {
	p := path
	if p[len(p)-1] != '/' {
		p += "/"
	}
	f := func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path
		if reqPath[len(reqPath)-1] != '/' {
			reqPath += "/"
		}
		if len(strings.Split(reqPath, "/"))-len(strings.Split(p, "/")) != 0 {
			notFound(w)
			return
		}
		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		cont.Set(w, r)
		reflect.ValueOf(cont).MethodByName(act).Call(nil)
	}

	Mux.HandleFunc(p, f)
}

func notFound(w http.ResponseWriter) {
	view, err := View("notFound")
	if err != nil {
		log.Fatal(err)
		return
	}
	err2 := view.Show(w)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}
