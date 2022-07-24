package product_service

import (
	"context"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/strutil"
	"github.com/olivere/elastic/v7"
	"shop-search-api/global"
	"strconv"
	"strings"
)

type Product struct {
	UserID   int64  `json:"user_id"`
	Keyword  string `json:"keyword"`
	New      *int   `json:"new"`
	Sales    string `json:"sales"`
	Price    string `json:"price"`
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

	//是否新品
	if s.New != nil {
		query.Must(elastic.NewTermQuery("is_new", s.New))
	}

	query.Should(shouldQuerys...)

	orders := make([]map[string]bool, 0)
	//价格排序
	if len(s.Price) > 0 {
		if strings.ToLower(s.Price) == "desc" {
			orders = append(orders, map[string]bool{"price": false})
		} else {
			orders = append(orders, map[string]bool{"price": true})
		}
	}

	//销量排序
	if len(s.Sales) > 0 {
		if strings.ToLower(s.Sales) == "desc" {
			orders = append(orders, map[string]bool{"sales": false})
		} else {
			orders = append(orders, map[string]bool{"sales": true})
		}
	}
	//默认按照相关度算分来排序
	orders = append(orders, map[string]bool{"_score": false})

	return global.ES.Query(context.Background(), global.ProductIndexName,
		nil, query, from, s.PageSize, es.WithEnableDSL(true),
		es.WithPreference(strconv.FormatInt(s.UserID, 10)),
		es.WithFetchSource(false), es.WithOrders(orders))
}
