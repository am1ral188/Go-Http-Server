package middleWares

import (
	"fmt"
	"net/http"
)

func ExampleMiddleWare(w http.ResponseWriter, r *http.Request, next func(http.ResponseWriter, *http.Request)) {
	fmt.Println("message from example middle ware : hello")
	next(w, r)
}
