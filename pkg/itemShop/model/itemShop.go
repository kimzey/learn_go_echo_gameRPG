package model

type (
	Item struct{
		ID          uint64    `json:"id"`
		Name        string    `json:"Name"`
		Description string    `json:"Description"`
		Picture     string    `json:"Picture"`
		Price       uint      `json:"Price"`
	}

	ItemFilter struct{
		Name string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=64"`
		Paginate
	}
	
	Paginate struct {
		Page int `query:"page" validate:"required,min=1"`
		Size int `query:"size" validate:"required,min=1,max=20"`
	}
	ItemResult struct{
		Items []*Item `json:"items"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct{
		Page int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}
)