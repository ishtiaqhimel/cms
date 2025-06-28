package utils

type Pagination struct {
	Page     int
	PageSize int
}

func (pg *Pagination) NextPage() *Pagination {
	pg.Page++
	return pg
}
