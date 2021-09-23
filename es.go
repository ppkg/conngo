package conngo

import (
	"log"

	"github.com/olivere/elastic/v7"
)

type Es struct {
	Url      string
	UserName string
	Password string
	client   *elastic.Client
}

// 获取一个现有的Client
func (es *Es) GetClient() *elastic.Client {
	if es.client == nil {
		es.client = es.NewClient()
	}
	return es.client
}

// 创建一个新的Client
func (es *Es) NewClient() *elastic.Client {
	if client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(es.Url),
		elastic.SetBasicAuth(es.UserName, es.Password),
	); err != nil {
		log.Println(err)
		return nil
	} else {
		return client
	}
}
