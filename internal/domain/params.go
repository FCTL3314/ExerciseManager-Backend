package domain

type FilterParams struct {
	Query interface{}
	Args  []interface{}
}

type PaginationParams struct {
	Limit  int
	Offset int
}

type OrderParams struct {
	Order string
}

type Params struct {
	Filter     FilterParams
	Pagination PaginationParams
	OrderParams
}
