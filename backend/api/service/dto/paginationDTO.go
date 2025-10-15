package dto

type PaginationDTO struct {
	Data        any  `json:"data"`
	CurrentPage uint `json:"current_page" example:"1"`
	TotalPages  uint `json:"total_pages" example:"10"`
	TotalItems  uint `json:"total_items" example:"100"`
}

func MakePaginationDTO[S any, D any](schemas []S, currentPage, totalPages, totalItems uint, toDTO func(S) D) (*PaginationDTO, error) {
	var DTOlist []D
	for _, item := range schemas {
		DTOlist = append(DTOlist, toDTO(item))
	}

	return &PaginationDTO{
		Data:        DTOlist,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
