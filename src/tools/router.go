package tools

type pro struct {
	Method     string
	Controller Controller
	Act        string
}
type action map[string]pro
type callBack func(*Router)

type Router struct {
	Path     string
	Children []*Router
	Actions  action
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
func (r *Router) GET(path string, cont Controller, act string) {
	r.Actions[path] = pro{"GET", cont, act}
}
func (r *Router) POST(path string, cont Controller, act string) {
	r.Actions[path] = pro{"POST", cont, act}
}
func (r *Router) init(obj *Router, parentPath string) {
	for s, f := range obj.Actions {
		Handle(parentPath+s, f.Controller, f.Act, f.Method)
	}

	p := parentPath + obj.Path
	for _, child := range obj.Children {
		r.init(child, p)
	}

}
func (r *Router) Init() {
	r.init(r, "")
}
