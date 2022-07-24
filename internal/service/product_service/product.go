package product_service

import (
	"context"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/strutil"
	"github.com/olivere/elastic/v7"
	"shop-search-api/global"
	"strconv"
)

type Product struct {
	UserID   int64  `json:"user_id"`
	Keyword  string `json:"keyword"`
	New      *int   `json:"new"`
	Sales    *int   `json:"sales"`
	Price    *int   `json:"price"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

func (s *Product) SearchProduct() (result *elastic.SearchResult, err error) {
	query := elastic.NewBoolQuery()
	from := s.PageNum * 20

	query.MinimumNumberShouldMatch(1)

	storeNameMatchPhreaseQuery := elastic.NewMatchPhraseQuery("store_name", s.Keyword).Boost(2).QueryName("storeNameMatchPhreaseQuery")
	storeNameMatchQuery := elastic.NewMatchPhraseQuery("store_name", s.Keyword).Boost(1).QueryName("storeNameMatchQuery")
	storeNamePinyinMatchPhreaseQuery := elastic.NewMatchPhraseQuery("store_name.pinyin", s.Keyword).Boost(0.7).QueryName("storeNamePinyinMatchPhreaseQuery")
	descMatchQuery := elastic.NewMatchPhraseQuery("desc", s.Keyword).Boost(0.5).QueryName("descMatchQuery")

	shouldQuerys := make([]elastic.Query, 0)
	shouldQuerys = append(shouldQuerys, storeNameMatchPhreaseQuery, storeNameMatchQuery, descMatchQuery)

	if strutil.IncludeLetter(s.Keyword) {
		shouldQuerys = append(shouldQuerys, storeNamePinyinMatchPhreaseQuery)
	}
	query.Should(shouldQuerys...)

	return global.ES.Query(context.Background(), global.ProductIndexName,
		nil, query, from, s.PageSize, es.WithEnableDSL(true),
		es.WithPreference(strconv.FormatInt(s.UserID, 10)), es.WithFetchSource(false))
}
