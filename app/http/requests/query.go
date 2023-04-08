package requests

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/boolh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/strh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/gofiber/fiber/v2"
)

type QueryFilter struct {
	// Column
	Column string
	// Operator, e.g. == or != or > or >= or < or <= or %% or --
	Operator string
	// Values
	Values []any
}

// Query represents a request query
type Query struct {
	// Filters is a list of where conditions
	// query = filters=name:%arfan%;gender:male|female;age:>=18;hobby:!=swimming;created_at:2020-2023 ,or via
	Filters []QueryFilter `json:"filters"`
	// OrderBys determines the order of the returned items
	// query = order_bys=name:asc;age:desc;
	OrderBys map[string]string `json:"order_bys"`
	// Withs determines the relation that will be loaded along the items
	// query = withs=users.user_profiles;users.users_settings
	Withs []string `json:"withs"`
	// Limit limits the returned items
	Limit int `json:"limit"`
	// Offset skip some items and then return the items after that offset/skip
	Offset int `json:"offset"`
}

// FromContext fills query from the given context
func (query *Query) FromContext(c *fiber.Ctx) error {
	errs := []error{
		query.setFiltersFromContext(c),
		query.setOrderBysFromContext(c),
		query.setWithsFromContext(c),
		query.setLimitOffsetFromContext(c),
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// AddFilter append the given QueryFilter to the Query.FIlters
func (query *Query) AddFilter(filters ...QueryFilter) {
	query.Filters = append(query.Filters, filters...)
}

// setFiltersFromContext will set Query.Filters from the the given context
func (query *Query) setFiltersFromContext(c *fiber.Ctx) (err error) {
	queryStr := c.Query("filters")

	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	var filters []QueryFilter
	filterStrings := strings.Split(queryStr, ";")
	for _, filterString := range filterStrings {
		syncronizer.WG().Add(1)
		go func(filterString string) {
			defer syncronizer.WG().Done()
			expression := regexp.MustCompile("^([\\w.]+):([=!<>%]{1,2})?([`]{1}[^`]+[`]{1})([|%-]{1,2})?(`{1}[^`]+`{1})?")
			filterArgs := expression.FindStringSubmatch(filterString)
			if len(filterArgs) == 0 {
				return
			}

			var filter QueryFilter
			filter.Column = filterArgs[1]
			filter.Operator = filterArgs[2]
			filter.Values = sliceh.Map(
				regexp.MustCompile("(`{1}[^`]+`{1})").FindAllString(filterString, -1),
				func(value string) any {
					return strings.Trim(value, "`")
				})
			secondOperator := filterArgs[4]

			if strings.Contains(filter.Operator, "!") {
				// if contains,then set as "not equal" operator
				filter.Operator = "!="
			} else if strings.Contains(filter.Operator, "%") || strings.Contains(secondOperator, "%") {
				// if contains,then set as "like" operator
				operator := boolh.Ternary(filter.Operator == "%", "%", ".")
				operator += boolh.Ternary(secondOperator == "%", "%", ".")
				filter.Operator = operator
			} else if sliceh.Contains([]string{">", ">=", "<", "<="}, filter.Operator) {
				// keep this empty
			} else if strings.Contains(secondOperator, "-") {
				// if contains,then set as "between" operator
				filter.Operator = "--"
			} else {
				// Default is "equal" operator
				filter.Operator = "=="
			}
			syncronizer.M().Lock()
			filters = append(filters, filter)
			syncronizer.M().Unlock()
		}(filterString)
	}
	syncronizer.WG().Wait()
	query.Filters = filters
	return
}

// setOrderBysFromContentex sets Query.OrderBy from the given contenxt
func (query *Query) setOrderBysFromContext(c *fiber.Ctx) (err error) {
	queryStr := c.Query("order_bys")

	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	orderBys := map[string]string{}
	orderByStrings := strings.Split(queryStr, ";")
	for _, orderByString := range orderByStrings {
		syncronizer.WG().Add(1)
		go func(orderByString string) {
			defer syncronizer.WG().Done()
			splitted := strings.Split(orderByString, ":")
			if len(splitted) != 2 {
				return
			}
			column := splitted[0]
			orderingType := splitted[1]
			orderBys[column] = orderingType
		}(orderByString)
	}
	syncronizer.WG().Wait()
	query.OrderBys = orderBys
	return
}

// setWithsFromContext will set Query.Withs from the the context
func (query *Query) setWithsFromContext(c *fiber.Ctx) (err error) {
	queryStr := c.Query("withs")
	var withs []string
	for _, with := range strings.Split(queryStr, ";") {
		with = strh.StrToDelimetedCamel(with, ".")
		withs = append(withs, with)
	}
	query.Withs = withs
	return
}

// setLimitOffsetFromContext
func (query *Query) setLimitOffsetFromContext(c *fiber.Ctx) (err error) {
	query.Limit = c.QueryInt("limit", c.QueryInt("per_page", 10))
	query.Offset = c.QueryInt(
		"offset",
		((c.QueryInt("page", 1) - 1) * query.Limit),
	)
	return
}

type QueryFilterExp struct {
	// Column
	Column string
	// Operator, e.g. == or != or > or >= or < or <= or %% or --
	Operator string
	// Value
	Value any
}

type QueryIncludeExp struct {
	Column  string
	Selects []string
}

type QuerySortExp struct {
	Column    string
	Direction string
}

type QueryExp struct {
	// querystring : select=column,column2,column3
	Selects []string
	// querystring : include=table:column1,column2;tabl2.table3;
	Includes []QueryIncludeExp
	// querystring : scope=hasComment,hasLike
	Scopes []string
	// querystring : withCount=table,table2
	Counts []string
	// querystring : sort=table.column:desc,table.column:rand
	Sorts []QuerySortExp
	// querystring : filter[column][operator]=value
	Filters []QueryFilterExp
}

// AddFilter append the given QueryFilter to the Query.FIlters
func (query *QueryExp) AddFilter(filters ...QueryFilterExp) {
	query.Filters = append(query.Filters, filters...)
}

func (input *QueryExp) FromContext(c *fiber.Ctx) (err error) {
	defer func() { // incase of error by unprocessable entity
		err = errorh.AnyToErrorOrNil(recover())
		if err != nil {
			err = fiber.ErrUnprocessableEntity
		}
	}()
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	fullUriStr := string(c.Request().URI().FullURI())
	url, err := url.Parse(fullUriStr)
	if err != nil {
		return
	}

	for key, values := range url.Query() {
		syncronizer.WG().Add(1)
		go func(key string, values []string) {
			defer syncronizer.WG().Done()
			if regexp.MustCompile("^selects?$").MatchString(key) {
				for _, values := range values {
					input.Selects = append(input.Selects, strings.Split(values, ",")...)
				}
			} else if regexp.MustCompile("^includes?$").MatchString(key) {
				for _, values := range values {
					for _, value := range strings.Split(values, ";") {
						var (
							splitted = strings.Split(value, ":")
							column   = splitted[0]
							selects  = []string{}
						)
						if len(splitted) > 1 && (splitted[1] != "") {
							selects = strings.Split(splitted[1], ",")
						}
						input.Includes = append(input.Includes, QueryIncludeExp{
							Column:  column,
							Selects: selects,
						})
					}
				}
			} else if regexp.MustCompile("^scopes?$").MatchString(key) {
				for _, values := range values {
					input.Scopes = append(input.Scopes, strings.Split(values, ",")...)
				}
			} else if regexp.MustCompile("^counts?$").MatchString(key) {
				for _, values := range values {
					input.Counts = append(input.Counts, strings.Split(values, ",")...)
				}
			} else if regexp.MustCompile("^sorts?$").MatchString(key) {
				for _, values := range values {
					for _, value := range strings.Split(values, ",") {
						splitted := strings.Split(value, ":")
						column := splitted[0]
						direction := splitted[1]
						input.Sorts = append(input.Sorts, QuerySortExp{
							Column:    column,
							Direction: direction,
						})
					}
				}
			} else if matcheds := regexp.MustCompile(
				fmt.Sprintf(`^filters?\[([\w.]+)\]\[(%s)\]`, strings.Join(enums.QueryFilterOperators(), "|")),
			).FindStringSubmatch(key); len(matcheds) != 0 {
				column := matcheds[1]
				operator := matcheds[2]
				input.AddFilter(QueryFilterExp{
					Column:   column,
					Operator: operator,
					Value:    values[0],
				})
			}
		}(key, values)
	}
	syncronizer.WG().Wait()

	return
}
