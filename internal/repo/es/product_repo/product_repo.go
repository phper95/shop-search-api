package product_repo

import "time"

type ProductIndex struct {
	Id          int64     `json:"id"`          //商品id
	StoreName   string    `json:"store_name"`  //商品名称
	StoreInfo   string    `json:"store_info"`  //商品简介
	Keyword     string    `json:"keyword"`     //关键字
	CateId      int       `json:"cate_id"`     //分类id
	Price       float64   `json:"price"`       //商品价格
	Sales       int32     `json:"sales"`       //销量
	Ficti       int32     `json:"ficti"`       //虚拟销量
	IsHot       int8      `json:"is_hot"`      //是否热卖 (0: 否，1：是)
	IsBenefit   int8      `json:"is_benefit"`  //是否优惠(0: 否，1：是)
	IsBest      int8      `json:"is_best"`     //是否精品(0: 否，1：是)
	IsNew       int8      `json:"is_new"`      //是否新品 (0: 否，1：是)
	Description string    `json:"description"` //产品描述
	IsPostage   int8      `json:"is_postage"`  //是否包邮 (0: 否，1：是)
	IsGood      int8      `json:"is_good"`     //是否优品推荐 (0: 否，1：是)
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}
