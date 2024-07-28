package Controller

import (
	"awesomeProject/src/cfg"
	"awesomeProject/src/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

type ApiController struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type response struct {
	Status bool
	Detail string
	Exists bool
}

func (r *ApiController) Set(w http.ResponseWriter, req *http.Request) {
	r.Response = w
	r.Request = req

}
func (r ApiController) Index() {
	io.WriteString(r.Response, "welcome to my landing page by api")
}

func (r ApiController) CardVerify() {
	r.Response.Header().Set("Content-Type", "application/json")
	var args []string
	args = append(args, r.Request.URL.Query().Get("user-id"), r.Request.URL.Query().Get("hash"))
	isValid := true
	for _, arg := range args {
		if arg == "" {
			isValid = false
		}
	}
	if isValid {
		matchString, err := regexp.MatchString("\\b[A-Fa-f0-9]{64}\\b", args[1])
		matchId, err2 := regexp.MatchString("\\b[a-zA-Z][a-zA-Z0-9]{5,15}\\b", args[0])
		if err2 != nil || err != nil {
			log.Fatal(err, err2)
			return
		}
		if matchString == false || matchId == false {
			res := response{Exists: false, Status: false, Detail: "Wrong Arguments"}
			jsonResponse, err3 := json.Marshal(res)
			if err2 != nil {
				fmt.Println(err3)
				return
			}
			_, err4 := io.WriteString(r.Response, string(jsonResponse))
			if err4 != nil {
				fmt.Println(err4)
				return
			}
			return
		}
		cMadel := model.CardModel{}
		db, errCon := cMadel.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)
		if errCon != nil {
			fmt.Println(errCon)
			return
		}
		err6 := db.AutoMigrate(cMadel.NewCard())
		if err6 != nil {
			fmt.Println(err6)
			return
		}

		row := db.Where("user_id=? AND hash=?", args[0], args[1])
		fmt.Println(row.RowsAffected)
		if row.Error != nil {
			fmt.Println("read error", row.Error)
			return
		}
		res := response{Exists: row.RowsAffected > 0, Status: true, Detail: fmt.Sprintf("does it exist: %t", row.RowsAffected > 0)}
		jsonResponse, err3 := json.Marshal(res)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		io.WriteString(r.Response, string(jsonResponse))
		return
	}
	res := response{Exists: false, Status: false, Detail: "nil or NO enough arguments"}
	jsonResponse, err3 := json.Marshal(res)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	io.WriteString(r.Response, string(jsonResponse))
	return
}

func (r ApiController) CardInsert() {
	r.Response.Header().Set("Content-Type", "application/json")
	var args []string
	args = append(args, r.Request.URL.Query().Get("user-id"), r.Request.URL.Query().Get("hash"))
	isValid := true
	for _, arg := range args {
		if arg == "" {
			isValid = false
		}
	}
	if isValid {
		matchString, err := regexp.MatchString("\\b[A-Fa-f0-9]{64}\\b", args[1])
		matchId, err2 := regexp.MatchString("\\b[a-zA-Z][a-zA-Z0-9]{5,15}\\b", args[0])
		if err2 != nil || err != nil {
			log.Fatal(err, err2)
			return
		}
		if matchString == false || matchId == false {
			res := response{Exists: false, Status: false, Detail: "Wrong Arguments"}
			jsonResponse, err3 := json.Marshal(res)
			if err2 != nil {
				fmt.Println(err3)
				return
			}
			_, err4 := io.WriteString(r.Response, string(jsonResponse))
			if err4 != nil {
				fmt.Println(err4)
				return
			}
			return
		}
		cMadel := model.CardModel{}
		db, errCon := cMadel.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)
		if errCon != nil {
			fmt.Println(errCon)
			return
		}
		err6 := db.AutoMigrate(cMadel.NewCard())
		if err6 != nil {
			fmt.Println(err6)
			return
		}
		row := db.Where("user_id=? AND hash=?", args)
		if row.Error != nil {
			fmt.Println(row.Error)
			return
		}
		res := response{Exists: row.RowsAffected == 0, Status: true, Detail: fmt.Sprintf("it aded : %t", row.RowsAffected == 0)}
		jsonResponse, err3 := json.Marshal(res)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		if row.RowsAffected < 1 {
			newRow := cMadel.NewCard()
			newRow.UserID = args[0]
			newRow.Hash = args[1]
			addedRow := db.Create(newRow)
			if addedRow.Error != nil {
				log.Fatal(addedRow.Error)
				return
			}
		}
		io.WriteString(r.Response, string(jsonResponse))
		return
	}
	res := response{Exists: false, Status: false, Detail: "nil or NO enough arguments"}
	jsonResponse, err3 := json.Marshal(res)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	io.WriteString(r.Response, string(jsonResponse))
	return
}

func NewApiController() *ApiController {
	return &ApiController{}
}
