package Post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"net/http"
	"github.com/cocoagaurav/httpHandler/model"
	)

func Fetchformhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Fetchform)
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	var (
		userid int
		title string
		description string
	)
	err:=database.Db.Ping()
	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		return
		}
	userpost:=&model.User{}
	err=json.NewDecoder(r.Body).Decode(userpost)
	if (err != nil) {
		w.WriteHeader(http.StatusNoContent)
	}
	data, err := database.Db.Query("select * from post where USERID=(select UID from user where name=?)", userpost.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	for data.Next() {
		err := data.Scan(&userid, &title, &description)
		if (err != nil) {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "USERID=%d \n\n title=%s \n\n description=%s \n\n\n\n ", userid, title, description)
	}
	defer data.Close()

	http.Redirect(w, r, "/fetchformhandler", 302)
}


