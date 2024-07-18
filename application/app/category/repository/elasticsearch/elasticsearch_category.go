package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"go-fiber-clean-architecture/application/domain"
	"strconv"
	"strings"
)

type ElasticCategoryRepository struct {
	Db *elastic.Client
}

func NewElasticCategoryRepository(db *elastic.Client) domain.CategoryRepository {
	return &ElasticCategoryRepository{db}
}

type ElasticStruct struct {
	
}

func (r *ElasticCategoryRepository) GetByID(ctx context.Context, id int64) (domain.Category, error) {

	query := `{ "query": { "match": {"id": "` +  strconv.FormatInt(id, 10) + `"} } }`
	search, _ := r.Db.Search(
		r.Db.Search.WithIndex("category"),
		r.Db.Search.WithBody(strings.NewReader(query)),
	)

	fmt.Println(search)

	panic("implement me")
	//get, err := r.Db.Get("category", "YaBr6I8BsJz1lu7l4Ytz")
	//if err != nil {
	//	panic(err)
	//}
	//// get response body
	//body, _ := io.ReadAll(get.Body)
	//var respond map[string]interface{}
	//// unmarshal body to struct
	//json.Unmarshal(body, &respond)
	//
	//
	//// decode struct to domain.Category
	//var category domain.Category
	//mapstructure.Decode(respond["_source"], &category)
	//return category, nil
}

func (r *ElasticCategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ElasticCategoryRepository) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	dataJson, err := json.Marshal(category)
	if err != nil {
		return domain.Category{}, err
	}

	data, err := r.Db.Index("category", bytes.NewReader(dataJson))

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	return category, nil
}

func (r *ElasticCategoryRepository) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ElasticCategoryRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *ElasticCategoryRepository) CreateAll(ctx context.Context, category []domain.Category) ([]domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ElasticCategoryRepository) CreateInBatches(ctx context.Context, category []domain.Category, size int) error {
	//TODO implement me
	panic("implement me")
}