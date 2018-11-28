package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/olivere/elastic"
	"log"
	"net/http"
)

func Fetchformhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Fetchform)
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	userpost := &model.User{}
	post := &model.Post{}
	err := json.NewDecoder(r.Body).Decode(userpost)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	esquery := elastic.NewTermQuery("id", userpost.Id)
	result, err := ElasticClient.Search("userpost").Index("userpost").Type("post").Query(esquery).Do(context.Background())
	if err != nil {
		log.Printf("error is: [%v]", err.Error())
	}
	for _, hit := range result.Hits.Hits {
		json.Unmarshal(*hit.Source, post)
		fmt.Fprintf(w, "USERID=%d \n\n title=%s \n\n description=%s \n\n\n\n ", userpost.Id, post.Title, post.Discription)
	}
	http.Redirect(w, r, "/fetchformhandler", 302)
}
