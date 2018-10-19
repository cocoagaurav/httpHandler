package Post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
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
	//fmt.Printf("uid is:%v", uid)

	newpost := &model.Post{}
	err = json.NewDecoder(r.Body).Decode(newpost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newpost.Id = uid.Id
	fmt.Printf("\n\npost id:%d \n post title:%s \n post disc:%s", newpost.Id, newpost.Title, newpost.Discription)

	jsonpost, err := json.Marshal(newpost)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("json data is:%s", string(jsonpost))

	MakeRabbitQ()

	Publish(jsonpost)

	ConsumeMssg()

	go func() {
		for msg := range Mssg {
			post := &model.Post{}
			data := bytes.NewReader(msg.Body)
			err := json.NewDecoder(data).Decode(post)
			if err != nil {
				log.Fatal(err)
				return
			}
			q, err := database.Db.Prepare("insert into post values(?,?,?)")
			defer q.Close()
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = q.Exec(post.Id, post.Title, post.Discription)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
	}()
	////fmt.Printf("newpost :[%+v]",newpost)

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
