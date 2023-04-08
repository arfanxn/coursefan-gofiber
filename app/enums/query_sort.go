package enums

// requests.QuerySort.Direction
const (
	QuerySortDirectionAscending  string = "asc"
	QuerySortDirectionDescending string = "desc"
	QuerySortDirectionRandom     string = "rand"
)

// QuerySortDirections returns slice of requests.QuerySort.Direction enums
func QuerySortDirections() []string {
	return []string{
		QuerySortDirectionAscending,
		QuerySortDirectionDescending,
		QuerySortDirectionRandom,
	}
}
