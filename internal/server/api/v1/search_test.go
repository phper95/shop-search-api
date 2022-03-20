package v1

import (
	"fmt"
	"gitee.com/phper95/pkg/httpclient"
	"net/http"
	"net/url"
	"shop-search-api/config"
	"shop-search-api/internal/pkg/sign"
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
	params.Add("keyword", "imooc")
	authorization, date, err := sign.New(ak, sk, ttl).Generate(ProductSearchUri, http.MethodGet, params)
	if err != nil {
		fmt.Println(err)
		return
	}
	headerAuth := httpclient.WithHeader(config.HeaderAuthField, authorization)
	headerAuthDate := httpclient.WithHeader(config.HeaderAuthDateField, date)
	r, e := httpclient.Get(ProductSearchHost+ProductSearchUri, params, headerAuth, headerAuthDate)
	fmt.Println(string(r), e)
}
