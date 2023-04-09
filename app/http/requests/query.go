package requests

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/boolh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/strh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

/*
 * Query
 */

type Query struct {
	// Column represents column to be included in the query
	// querystring/formdata : column=id,name,email
	Columns []string
	// Inlcludes represents relations to be included in the query
	// querystring/formdata : include[users]=id,name,email
	Includes []QueryInclude
	// Scopes represents query scope to be included in the query
	// querystring/formdata : scope=hasFriends,hasEmployees
	Scopes []string
	// Aggregates represents aggregates to be included in the query
	// querystring/formdata : aggregate[users]=count or aggregate[wallet.balance]=sum
	Aggregates []QueryAggregate
	// Sorts represents the sort order to be included in the query
	// querystring/formdata : sort[users.created_at]=desc
	Sorts []QuerySort
	// Filters represents the filter to be included in the query
	// querystring/formdata : filter[users.name][eq]=Jack
	Filters []QueryFilter
	// Limit represents the limit items to be returned in the query
	Limit null.Int
	// Offset represents the total items to be offseted in the query
	Offset int
}

// AddFilter append the given QueryFilter to the Query.FIlters
func (query *Query) AddFilter(filters ...QueryFilter) {
	query.Filters = append(query.Filters, filters...)
}

func (input *Query) FromContext(c *fiber.Ctx) (err error) {
	defer func() { // incase of error by unprocessable entity
		err = errorh.AnyToErrorOrNil(recover())
		if err != nil {
			err = fiber.ErrUnprocessableEntity
		}
	}()
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	url, err := url.Parse(ctxh.GetFullURIString(c))
	if err != nil {
		return
	}

	urlQueries := url.Query()

	// Pagination Limit / Per Page
	if limit, err := strconv.ParseInt(urlQueries.Get("limit"), 10, 64); err == nil {
		input.Limit = null.NewInt(limit, true)
	} else if perPage, err := strconv.ParseInt(urlQueries.Get("per_page"), 10, 64); err == nil {
		input.Limit = null.NewInt(perPage, true)
	}
	// Pagination Offset / Page
	if offset, err := strconv.ParseInt(urlQueries.Get("offset"), 10, 64); err == nil {
		input.Offset = int(offset)
	} else if page, err := strconv.ParseInt(urlQueries.Get("page"), 10, 64); err == nil {
		input.Offset = (int(page) - 1) * int(input.Limit.Int64)
	}

	for key, values := range urlQueries {
		syncronizer.WG().Add(1)
		go func(key string, values []string) {
			defer syncronizer.WG().Done()
			if err := syncronizer.Err(); err != nil {
				return
			}
			syncronizer.RWM().Lock()
			defer syncronizer.RWM().Unlock()
			if regexp.MustCompile(`^(selects?|columns?)$`).MatchString(key) {
				// Column selection
				input.Columns = sliceh.Merge(
					input.Columns,
					boolh.Ternary(len(values) != 0, strings.Split(values[0], ","), []string{"*"}),
				)
			} else if matcheds := regexp.MustCompile(`^(includes?\[([\w.]+)\])$`).
				// Relation loader
				FindStringSubmatch(key); len(matcheds) != 0 {
				var (
					limitKey     = matcheds[1] + "[limit]"
					nullIntLimit = null.NewInt(0, false)
					relation     = matcheds[2]
				)
				if values, ok := urlQueries[limitKey]; ok && (len(values) != 0) && (values[0] != "") {
					if limit, err := strconv.ParseInt(values[0], 10, 64); err == nil {
						nullIntLimit = null.NewInt(limit, true)
					}
				}
				input.Includes = append(input.Includes, QueryInclude{
					Relation: relation,
					Columns: boolh.Ternary(
						(len(values) != 0) && (values[0] != ""), strings.Split(values[0], ","), []string{"*"},
					),
					Limit: nullIntLimit,
				})
			} else if regexp.MustCompile("^scopes?$").MatchString(key) {
				// Query scopes
				input.Scopes = sliceh.Merge(input.Scopes, strings.Split(values[0], ","))
			} else if matcheds := regexp.MustCompile(`^aggregates?\[([\w.]+)\]$`).
				// Query Aggregates
				FindStringSubmatch(key); len(matcheds) != 0 {
				column := matcheds[1]
				input.Aggregates = append(input.Aggregates, QueryAggregate{
					Column: column,
					Name:   values[0],
				})
			} else if matcheds := regexp.MustCompile(`^sorts?\[([\w.]+)\]$`).
				// Query Sorts
				FindStringSubmatch(key); len(matcheds) != 0 {
				column := matcheds[1]
				direction := values[0]
				input.Sorts = append(input.Sorts, QuerySort{
					Column:    column,
					Direction: direction,
				})
			} else if matcheds := regexp.MustCompile(
				fmt.Sprintf(`^filters?\[([\w.]+)\]\[(%s)\]$`, strings.Join(enums.QueryFilterOperators(), "|")),
			).FindStringSubmatch(key); len(matcheds) != 0 {
				// Query Filter
				column := matcheds[1]
				operator := matcheds[2]
				input.AddFilter(QueryFilter{
					Column:   column,
					Operator: operator,
					Values:   sliceh.Map(values, strh.StrToAny),
				})
			}
		}(key, values)
	}
	syncronizer.WG().Wait()
	if err = syncronizer.Err(); err != nil {
		return
	}

	return
}

/*
 * Query Filter
 */

type QueryFilter struct {
	// Column
	Column string
	// Operator, e.g. == or != or > or >= or < or <= or %% or --
	Operator string
	// Values
	Values []any
}

// TableName returns the table name of the QueryFilter.Column
func (filter QueryFilter) TableName() string {
	return strings.Split(filter.Column, ".")[0]
}

// CastValues casts data on QueryFilter.Values to the appropriate types, this function also returns the casted QueryFilter.Values
func (filter *QueryFilter) CastValues() []any {
	filter.Values = sliceh.Map(filter.Values, func(value any) any {
		valueStr := strh.AnyToStr(value)
		if value, err := strconv.ParseFloat(valueStr, 10); err == nil {
			return value
		} else if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		} else if value, err := time.Parse(time.RFC3339, valueStr); err == nil {
			return value
		} else {
			return valueStr
		}
	})
	return filter.Values
}

/*
 * Query Include
 */

type QueryInclude struct {
	// Relationble represents the table(s) will be included/joined in the query
	Relation string
	// Column represents column to be included in the query
	Columns []string
	// Limit represents the number of rows to be returned in the query
	Limit null.Int
}

// TableNames returns the relation table names of the QueryInclude.Table
func (include QueryInclude) TableNames() []string {
	return strings.Split(include.Relation, ".")
}

/*
 * Query Sort
 */

type QuerySort struct {
	// Column represents the column being sorted
	Column string
	// Direction represents the sorting direction
	Direction string
}

// TableName returns the table name of the QuerySort.Column
func (filter QuerySort) TableName() string {
	return strings.Split(filter.Column, ".")[0]
}

/*
 * Query Aggregate
 */

type QueryAggregate struct {
	// Column store the table or column being aggregated
	Column string
	// Name is the name of the aggregation e.g: count, sum, avg, min, max, etc.
	Name string
}
