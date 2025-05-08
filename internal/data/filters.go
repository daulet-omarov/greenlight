package data

import (
	"github.com/daulet-omarov/greenlight/internal/validator"
	"strings"
)

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

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}
