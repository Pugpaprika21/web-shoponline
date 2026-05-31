package dto

// PaginationRequest holds pagination parameters
type PaginationRequest struct {
	Page     int `query:"page" json:"page"`
	PageSize int `query:"page_size" json:"page_size"`
}

// PaginationResponse holds pagination metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

func (p *PaginationRequest) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		p.PageSize = 10
	}
}

func (p *PaginationRequest) Offset() int {
	return (p.Page - 1) * p.PageSize
}
