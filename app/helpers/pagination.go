package helpers

import "gorm.io/gorm"

type Pagination struct {
	Page      int    `form:"page" query:"page" json:"page"`
	Limit     int    `form:"limit" query:"limit" json:"limit"`
	KeySearch string `json:"key_search"`
	OrderBy   string `json:"order_by"`
	Sort      string `json:"sort"`
	Schema    string `json:"schema"`
}

func (p *Pagination) GetOffset() int {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	return (p.Page - 1) * p.Limit
}

func Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
