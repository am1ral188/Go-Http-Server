package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViewFile struct {
	fileData string
}

func (r ViewFile) Show(w http.ResponseWriter) error {
	_, err := io.WriteString(w, r.fileData)
	return err
}
func View(name string) (ViewFile, error) {

	info, err2 := os.Stat("./src/view/" + name + ".html")
	if err2 != nil {
		return ViewFile{}, err2
	}
	f, err := os.Open("./src/view/" + name + ".html")
	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		fmt.Println(err)
		return ViewFile{}, err
	} else {
		return ViewFile{fileData: string(data)}, nil
	}

}
