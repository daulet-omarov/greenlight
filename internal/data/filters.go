package data

import "github.com/daulet-omarov/greenlight/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "page must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "page must be less than or equal to 10_000_000")
	v.Check(f.PageSize > 0, "page_size", "page_size must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "page_size must be less than or equal to 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
