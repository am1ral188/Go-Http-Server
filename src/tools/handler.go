package tools

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Controller interface {
	Set(http.ResponseWriter, *http.Request)
}

var Mux = http.NewServeMux()

type middleWare struct {
	RW         http.ResponseWriter
	Req        *http.Request
	Cont       Controller
	Action     string
	MiddleWare func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
}

func (m *middleWare) Set(w http.ResponseWriter, r *http.Request) {
	m.RW = w
	m.Req = r
}
func (m *middleWare) Index(args ...interface{}) {
	f := func(w http.ResponseWriter, r *http.Request) {
		m.Cont.Set(w, r)
		params := []reflect.Value{}
		if len(args) == 0 {
			params = nil
		}
		for _, arg := range args {
			params = append(params, reflect.ValueOf(arg))
		}
		reflect.ValueOf(m.Cont).MethodByName(m.Action).Call(params)
	}
	m.MiddleWare(m.RW, m.Req, f)
}

func Handle(path string, cont Controller, act string, method string, w http.ResponseWriter, r *http.Request, handleSubPath bool) bool {
	if r.Method != method {
		return false
	}
	p := path
	if p[len(p)-1] != '/' {
		p += "/"
	}
	reqP := r.URL.Path
	if reqP[len(reqP)-1] != '/' {
		reqP += "/"
	}
	pSplit := strings.Split(p, "/")
	reqPathSplit := strings.Split(reqP, "/")
	if len(reqPathSplit)-len(pSplit) < 0 {
		return false
	}
	if !handleSubPath && len(reqPathSplit)-len(pSplit) != 0 {
		return false
	}
	var args []string
	finalP := pSplit
	for i, s := range pSplit {
		if len(s) < 3 {
			continue
		}
		if s[0] == '{' && s[len(s)-1] == '}' {
			nameAndRegex := strings.Split(s[1:len(s)-1], ":")
			matchString, err := regexp.MatchString(nameAndRegex[1], reqPathSplit[i])
			if err != nil {
				return false
			}
			if matchString {
				finalP[i] = reqPathSplit[i]
				args = append(args, reqPathSplit[i])
			} else {
				return false
			}
		}
	}
	if strings.Join(reqPathSplit, "/") == strings.Join(finalP, "/") || (handleSubPath && strings.Contains(strings.Join(reqPathSplit, "/"), strings.Join(finalP, "/"))) {
		cont.Set(w, r)
		if len(args) == 0 {
			reflect.ValueOf(cont).MethodByName(act).Call(nil)
		} else {
			params := []reflect.Value{}
			for _, arg := range args {
				params = append(params, reflect.ValueOf(arg))
			}
			reflect.ValueOf(cont).MethodByName(act).Call(params)
		}
		return true
	}
	return false
}

func HandleWithMiddleWare(path string, cont Controller, act string, method string, w http.ResponseWriter, r *http.Request, m func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request)), subPathHandle bool) bool {
	return Handle(path, &middleWare{MiddleWare: m, Cont: cont, Action: act}, "Index", method, w, r, subPathHandle)
}
