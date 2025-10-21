package common

type Pagination struct {
	Page      int  `json:"page"`
	Limit     int  `json:"limit"`
	PerPage   int  `json:"per_page"`
	TotalPage *int `json:"total_page,omitempty"`
	Total     *int `json:"total,omitempty"`
}

func (p *Pagination) WrapTotalPage(tp int) *Pagination {
	p.TotalPage = &tp
	return p
}

func (p *Pagination) WrapTotal(tt int) *Pagination {
	p.Total = &tt
	return p
}