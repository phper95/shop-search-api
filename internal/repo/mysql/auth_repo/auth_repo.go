package auth_repo

import (
	"fmt"
	"gitee.com/phper95/pkg/errors"
	"gorm.io/gorm"
	"time"
)

const (
	IsUsedYES = 1  // 启用
	IsUsedNo  = -1 // 禁用
)

type Auth struct {
	Id                int32     // 主键
	BusinessKey       string    // 调用方key
	BusinessSecret    string    // 调用方secret
	BusinessDeveloper string    // 调用方对接人
	Remark            string    // 备注
	IsUsed            int32     // 是否启用 1:是  -1:否
	IsDeleted         int32     // 是否删除 1:是  -1:否
	CreatedAt         time.Time `gorm:"time"` // 创建时间
	CreatedUser       string    // 创建人
	UpdatedAt         time.Time `gorm:"time"` // 更新时间
	UpdatedUser       string    // 更新人
}

func NewModel() *Auth {
	return new(Auth)
}

func NewQueryBuilder() *authQueryBuilder {
	return new(authQueryBuilder)
}

func (t *Auth) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type authQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *authQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *authQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Auth{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *authQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Auth{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *authQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Auth{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *authQueryBuilder) First(db *gorm.DB) (*Auth, error) {
	ret := &Auth{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *authQueryBuilder) QueryOne(db *gorm.DB) (*Auth, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *authQueryBuilder) QueryAll(db *gorm.DB) ([]*Auth, error) {
	var ret []*Auth
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *authQueryBuilder) Limit(limit int) *authQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *authQueryBuilder) Offset(offset int) *authQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *authQueryBuilder) WhereId(predicate string, value int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIdIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIdNotIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderById(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *authQueryBuilder) WhereBusinessKey(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessKeyIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessKeyNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByBusinessKey(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_key "+order)
	return qb
}

func (qb *authQueryBuilder) WhereBusinessSecret(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessSecretIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessSecretNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByBusinessSecret(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_secret "+order)
	return qb
}

func (qb *authQueryBuilder) WhereBusinessDeveloper(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessDeveloperIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereBusinessDeveloperNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByBusinessDeveloper(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_developer "+order)
	return qb
}

func (qb *authQueryBuilder) WhereRemark(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereRemarkIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereRemarkNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByRemark(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "remark "+order)
	return qb
}

func (qb *authQueryBuilder) WhereIsUsed(predicate string, value int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIsUsedIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIsUsedNotIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByIsUsed(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *authQueryBuilder) WhereIsDeleted(predicate string, value int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIsDeletedIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereIsDeletedNotIn(value []int32) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByIsDeleted(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *authQueryBuilder) WhereCreatedAt(predicate string, value time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereCreatedAtIn(value []time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByCreatedAt(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *authQueryBuilder) WhereCreatedUser(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereCreatedUserIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereCreatedUserNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByCreatedUser(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedAt(predicate string, value time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedAtIn(value []time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByUpdatedAt(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedUser(predicate, value string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", predicate),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedUserIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) WhereUpdatedUserNotIn(value []string) *authQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authQueryBuilder) OrderByUpdatedUser(asc bool) *authQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
