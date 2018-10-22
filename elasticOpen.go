package main

import (
	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
)

func ElasticConn() {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
}
