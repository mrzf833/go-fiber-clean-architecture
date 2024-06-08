package utils

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"go-fiber-clean-architecture/application/config"
)

func GetESClient() (*elastic.Client, error) {
	client, err :=  elastic.NewClient(elastic.Config{
		Addresses: []string{"http://" + config.ElasticHost + ":" + config.ElasticPort},
		Username: config.ElasticUser,
		Password: config.ElasticPassword,
	})

	fmt.Println("ES initialized...")
	return client, err
}

func ConnectElasticsearch()  {
	client, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing ES client")
		panic(err)
	}
	config.ElasticDb = client
}