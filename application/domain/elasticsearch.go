package domain

type ElastichSearchResponse struct {
	Shards   ShardResponse `json:"_shards"`
	Hits     HitsResponse  `json:"hits"`
	TimedOut bool          `json:"timed_out"`
	Took     int           `json:"took"`
}

type ShardResponse struct {
	Failed     int `json:"failed"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
	Skipped    int `json:"skipped"`
}

type HitsResponse struct {
	Hits     []map[string]interface{} `json:"hits"`
	Total    HitsTotalResponse        `json:"total"`
	MaxScore float64                  `json:"max_score"`
}

type HitsTotalResponse struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}

type DataResponse struct {
	Id     string                 `json:"_id" mapstructure:"_id"`
	Index  string                 `json:"_index" mapstructure:"_index"`
	Source map[string]interface{} `json:"_source" mapstructure:"_source"`
	Score  int                    `json:"_score" mapstructure:"_score"`
}
