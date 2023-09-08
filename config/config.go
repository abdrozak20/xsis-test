package config

import (
	"math"
	"time"
)

var defaultlimit int = 10

type PaginationRequest struct {
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
	Search string `query:"search"`
}

type GeneralFilters struct {
	Limit  int
	Page   int
	Search string
}

type Meta struct {
	Limit      int   `json:"per_page"`
	Offset     int   `json:"-"`
	Page       int   `json:"current_page"`
	TotalRows  int64 `json:"total"`
	TotalPages int   `json:"last_page"`
}

func (m *Meta) SetOffset() {
	m.Offset = (m.Page - 1) * m.Limit
}

func (m *Meta) SetLimit(limit int) {
	m.Limit = limit
	if m.Limit == 0 {
		m.Limit = defaultlimit
	}
}

func (m *Meta) SetPage(page int) {
	m.Page = page
	if m.Page == 0 {
		m.Page = 1
	}
}
func (m *Meta) GetLimit() int {
	return m.Limit
}

func (m *Meta) GetPage() int {
	return m.Page
}

func (m *Meta) GetMeta(totalRows int64, totalCurrentRow int) {
	m.SetOffset()

	m.TotalRows = totalRows
	m.Page = m.GetPage()
	m.Limit = m.GetLimit()
	m.TotalPages = int(math.Ceil(float64(m.TotalRows) / float64(m.Limit)))
}

func TimeMillisecond() int64 {
	now := time.Now()
	unixMilli := now.UnixMilli()

	return unixMilli
}
