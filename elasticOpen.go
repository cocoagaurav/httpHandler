package main

import (
	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
	"time"
)

func ElasticConn() {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))
	if err != nil {
		log.Printf("err=[%v]", err)
		time.Sleep(5 * time.Second)
		ElasticConn()
		//log.Fatal(err)
	}
	log.Printf("elastic search is connected..................")
}
