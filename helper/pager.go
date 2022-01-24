package helper

import "math"

type Pager struct {
	Total       int         `json:"total"`
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	PerPage     int         `json:"perPage"`
}

func (p *Pager) RunPage(page int, size int, total int, data interface{}) *Pager {
	p.Total = total
	p.Data = data
	p.CurrentPage = page
	p.PerPage = size
	p.LastPage = int(math.Ceil(float64(total / size)))
	return p
}
