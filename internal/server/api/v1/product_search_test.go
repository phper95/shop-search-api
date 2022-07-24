package v1

import (
	"fmt"
	"gitee.com/phper95/pkg/httpclient"
	"gitee.com/phper95/pkg/sign"
	"net/http"
	"net/url"
	"shop-search-api/config"
	"testing"
	"time"
)

const ProductSearchHost = "http://127.0.0.1:9090"
const ProductSearchUri = "/api/v1/product-search"

var (
	ak  = "AK100523687952"
	sk  = "W1WTYvJpfeH1YpUjTpeFbEx^DnpQ&35L"
	ttl = time.Minute * 3
)

func TestProductSearch(t *testing.T) {
	params := url.Values{}
	params.Add("userid", "1")
	params.Add("keyword", "手机")
	params.Add("page_num", "1")
	params.Add("page_size", "10")
	authorization, date, err := sign.New(ak, sk, ttl).Generate(ProductSearchUri, http.MethodGet, params)
	if err != nil {
		fmt.Println(err)
		return
	}
	headerAuth := httpclient.WithHeader(config.HeaderAuthField, authorization)
	headerAuthDate := httpclient.WithHeader(config.HeaderAuthDateField, date)
	c, r, e := httpclient.Get(ProductSearchHost+ProductSearchUri, params, headerAuth, headerAuthDate)
	fmt.Println(c, string(r), e)
}
