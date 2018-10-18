package Post

import (
	"fmt"
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"net/http"
)

func AfterLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPages.InternalPage)
}
func Posthandler(w http.ResponseWriter, r *http.Request) {
	err := database.Db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	c, err := r.Cookie("sessiontoken")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	uid := database.UserCache[c.Value]

	//newpost:=&model.Post{}
	//err=json.NewDecoder(r.Body).Decode(newpost)
	//if(err!=nil){
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	////fmt.Printf("newpost :[%+v]",newpost)
	//
	//q, err := database.Db.Prepare("insert into post values(?,?,?)")
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//defer q.Close()
	//_, err = q.Exec(uid.Id, newpost.Title, newpost.Discription)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	http.Redirect(w, r, "/success", 302)
}
