package common

type Params struct {
	Page  int64 `form:"page" example:"1" format:"int"`
	Limit int64 `form:"limit" example:"10" format:"int"`
	St    int   `form:"st" example:"30" format:"int"`
	En    int   `form:"en" example:"30" format:"int"`
}

type Paginator struct {
	Page  int64 `json:"page" example:"1" format:"int"`
	Limit int64 `json:"limit" example:"30" format:"int"`
	Total int64 `json:"total" example:"100" format:"int"`
	St    int64 `json:"st" example:"30" format:"int"`
	En    int64 `json:"en" example:"30" format:"int"`
}

func (cp *Params) Check() {
	if cp.Limit == 0 {
		cp.Limit = 20
	}
	if cp.Page == 0 {
		cp.Page = 1
	}
}

func (paginator *Paginator) SetTotal(total int64) {
	paginator.Total = total
}

func (paginator *Paginator) GetStart() int64 {
	start := (paginator.Page - 1) * paginator.Limit
	if start < 0 {
		start = 0
	}
	return start
}

func GenPaginator(commonParam Params) Paginator {
	commonParam.Check()
	page := commonParam.Page
	limit := commonParam.Limit

	return Paginator{
		Page:  page,
		Limit: limit,
	}
}
