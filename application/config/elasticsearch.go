package config

import (
	elastic "github.com/elastic/go-elasticsearch/v8"
	"go-fiber-clean-architecture/application/utils/helper2"
)

var (
	ElasticDb *elastic.Client
	ElasticPort = helper2.GetEnv("ELASTIC_PORT", "9200")
	ElasticHost = helper2.GetEnv("ELASTIC_HOST", "localhost")
	ElasticUser = helper2.GetEnv("ELASTIC_USER", "")
	ElasticPassword = helper2.GetEnv("ELASTIC_PASSWORD", "")
)