package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"go-fiber-clean-architecture/application/utils"
	"io"
	"strconv"
	"strings"
)

type elasticProductRepository struct {
	Db *elastic.Client
}

func NewElasticProductRepository(db *elastic.Client) domain.ProductRepository {
	return &elasticProductRepository{db}
}

func (r elasticProductRepository) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	query := `{ "query": { "match": {"id": "` +  strconv.FormatInt(id, 10) + `"} } }`
	search, _ := r.Db.Search(
		r.Db.Search.WithIndex("product"),
		r.Db.Search.WithBody(strings.NewReader(query)),
	)

	body, _ := io.ReadAll(search.Body)
	var respond domain.ElastichSearchResponse
	// unmarshal body to struct
	json.Unmarshal(body, &respond)

	if len(respond.Hits.Hits) == 0 {
		return domain.Product{}, exception.ErrNotFound
	}

	var data domain.DataResponse
	mapstructure.Decode(respond.Hits.Hits[0], &data)

	var product domain.Product
	utils.MapStructureDecode(data.Source, &product)

	return product, nil
}

func (r elasticProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	query := `{ "query": { "match_all": {} } }`
	search, _ := r.Db.Search(
		r.Db.Search.WithIndex("product"),
		r.Db.Search.WithBody(strings.NewReader(query)),
	)

	body, _ := io.ReadAll(search.Body)

	var respond domain.ElastichSearchResponse
	// unmarshal body to struct
	json.Unmarshal(body, &respond)

	for _, hitsRespond := range respond.Hits.Hits {
		var product domain.Product
		var data domain.DataResponse
		mapstructure.Decode(hitsRespond, &data)
		utils.MapStructureDecode(data.Source, &product)

		products = append(products, product)
	}

	return products, nil
}

func (r elasticProductRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	dataJson, err := json.Marshal(product)
	if err != nil {
		return domain.Product{}, err
	}

	_, err = r.Db.Index("product", bytes.NewReader(dataJson))

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r elasticProductRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r elasticProductRepository) Delete(ctx context.Context, id int64) error {
	query := `{ "query": { "match": {"id": "` +  strconv.FormatInt(id, 10) + `"} } }`
	search, _ := r.Db.Search(
		r.Db.Search.WithIndex("product"),
		r.Db.Search.WithBody(strings.NewReader(query)),
	)

	body, _ := io.ReadAll(search.Body)
	var respond domain.ElastichSearchResponse
	// unmarshal body to struct
	json.Unmarshal(body, &respond)

	if len(respond.Hits.Hits) == 0 {
		return exception.ErrNotFound
	}

	var data domain.DataResponse
	mapstructure.Decode(respond.Hits.Hits[0], &data)

	_, err := r.Db.Delete("product", data.Id)

	return err
}
