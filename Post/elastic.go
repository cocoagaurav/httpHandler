package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
)

func ElasticOpen() {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
}
func Datainsert(row string) {
	_, ierr := ElasticClient.Index().Index("userpost").Type("post").BodyJson(row).Do(context.Background())
	if ierr != nil {
		fmt.Printf("datainsert error:=%s", ierr.Error())
	}
}
