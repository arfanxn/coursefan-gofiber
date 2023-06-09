package resources

import (
	"net/url"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

type Pagination[T any] struct {
	// The total items in current pagination
	Total int `json:"total"`

	// Page information
	PerPage     int      `json:"per_page"`
	FirstPage   int      `json:"first_page"`
	PrevPage    null.Int `json:"prev_page"`
	CurrentPage int64    `json:"current_page"`
	NextPage    null.Int `json:"next_page"`
	LastPage    null.Int `json:"last_page,omitempty"`

	// Pagination url navigations
	FirstPageUrl   string      `json:"first_page_url"`
	PrevPageUrl    null.String `json:"prev_page_url"`
	CurrentPageUrl string      `json:"current_page_url"`
	NextPageUrl    null.String `json:"next_page_url"`
	LastPageUrl    null.String `json:"last_page_url,omitempty"`

	// Pagination items/data
	Items []T `json:"items"`
}

// SetPage set the pagination page related fields,
// Just give null on lastPage parameter if the pagination does not have last page information
func (pagination *Pagination[T]) SetPage(perPage int, currentPage int64, lastPage null.Int) {
	pagination.PerPage = perPage // set per page

	pagination.FirstPage = 1 // set first page

	// only set if prev page is greater than or equal to 1
	if prevPage := currentPage - 1; prevPage >= 1 {
		pagination.PrevPage = null.NewInt(prevPage, true) // set prev page
	}

	pagination.CurrentPage = currentPage // set current page

	nextPage := currentPage + 1

	// only set if last page is valid and next page is less than or equal last page
	if (lastPage.Valid) && (nextPage <= lastPage.Int64) {
		pagination.NextPage = null.NewInt(currentPage+1, true) // set next page
	} else if !lastPage.Valid { // if last page is not valid then set next page
		pagination.NextPage = null.NewInt(nextPage, true)
	}

	// only set if last page is valid and last page is greater than or equal current page
	if (lastPage.Valid) && (lastPage.Int64 >= currentPage) {
		pagination.LastPage = null.NewInt(lastPage.Int64, true) // set last page
	}
}

// SetPageFromOffsetLimit set the pagination page related fields from the given offset and limit,
func (pagination *Pagination[T]) SetPageFromOffsetLimit(offset int64, limit int) {
	pagination.SetPage(limit, (offset/int64(limit))+1, null.NewInt(0, false))
}

// SetItems sets the pagination.Items and pagination.Total
func (pagination *Pagination[T]) SetItems(items []T) {
	pagination.Total = len(items)
	pagination.Items = items
}

// SetURL set the pagination url related fields
func (pagination *Pagination[T]) SetURL(urlStruct *url.URL) error {
	if urlStruct == nil {
		return nil
	}

	page, err := int64(0), error(nil)
	if pageStr := urlStruct.Query().Get("page"); pageStr != "" {
		if page, err = strconv.ParseInt(pageStr, 10, 64); err != nil {
			return err
		}
	} else {
		page = 1
	}

	pagination.CurrentPage = page

	// Set the pagination first page url
	firstPageUrlStruct := *urlStruct
	firstPageUrlQueries := firstPageUrlStruct.Query()
	firstPageUrlQueries.Set("page", strconv.Itoa(1))
	firstPageUrlStruct.RawQuery = firstPageUrlQueries.Encode()
	pagination.FirstPageUrl = firstPageUrlStruct.String()

	pagination.CurrentPageUrl = urlStruct.String() // set current page url

	// Set the pagination prev page url
	prevPage := pagination.CurrentPage - 1
	if prevPage >= 1 { // only set the pagination prev page url if prev page is gte 1
		prevPageUrlStruct := *urlStruct
		prevPageUrlQueries := prevPageUrlStruct.Query()
		prevPageUrlQueries.Set("page", strconv.Itoa(int(prevPage)))
		prevPageUrlStruct.RawQuery = prevPageUrlQueries.Encode()
		pagination.PrevPageUrl = null.NewString(prevPageUrlStruct.String(), true)
	}

	// Set the pagination next page url
	// only set the pagination next page url if pagination items are not empty
	if len(pagination.Items) > 1 {
		nextPage := pagination.CurrentPage + 1
		nextPageUrlStruct := *urlStruct
		nextPageUrlQueries := nextPageUrlStruct.Query()
		nextPageUrlQueries.Set("page", strconv.Itoa(int(nextPage)))
		nextPageUrlStruct.RawQuery = nextPageUrlQueries.Encode()
		pagination.NextPageUrl = null.NewString(nextPageUrlStruct.String(), true)
	}

	// Set the pagination last page url
	// only set the pagination last page url if last page is valid and gte current page
	if pagination.LastPage.Valid && (pagination.LastPage.Int64 >= pagination.CurrentPage) {
		lastPageUrlStruct := *urlStruct
		lastPageUrlQueries := lastPageUrlStruct.Query()
		lastPageUrlQueries.Set("page", strconv.Itoa(int(pagination.LastPage.Int64)))
		lastPageUrlStruct.RawQuery = lastPageUrlQueries.Encode()
		pagination.LastPageUrl = null.NewString(lastPageUrlStruct.String(), true)
	}
	return nil
}
