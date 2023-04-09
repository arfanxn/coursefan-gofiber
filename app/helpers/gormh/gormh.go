// GORM helper package`
package gormh

import (
	"fmt"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/strh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gorm.io/gorm"
)

// BuildFromRequestQuery builds a gorm query from the given requests.Query
func BuildFromRequestQuery(tx *gorm.DB, query requests.Query) *gorm.DB {
	mainModel := models.User{}
	tx = tx.Table(mainModel.TableName())
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	for _, filter := range query.Filters {
		if filter.TableName() != mainModel.TableName() {
			continue
		}
		filter.CastValues()
		tx = buildFromRequestQueryFilter(tx, filter)
	}
	for _, include := range query.Includes {
		tx.Preload(strh.StrToDelimetedCamel(include.Relation, "."), func(tx *gorm.DB) *gorm.DB {
			relationTableNames := include.TableNames()
			lastRelationTableName := sliceh.Last(relationTableNames)
			for _, filter := range query.Filters {
				if filter.TableName() != lastRelationTableName {
					continue
				}
				tx = buildFromRequestQueryFilter(tx, filter)
			}
			for _, sort := range query.Sorts {
				if sort.TableName() != lastRelationTableName {
					continue
				}
				tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Direction))
			}
			if include.Limit.Valid {
				tx = tx.Limit(int(include.Limit.Int64))
			} else if query.Limit.Valid {
				tx = tx.Limit(int(query.Limit.Int64))
			}
			return tx
		})
	}
	for _, sort := range query.Sorts {
		if sort.TableName() != mainModel.TableName() {
			continue
		}
		tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Direction))
	}
	if query.Limit.Valid {
		tx = tx.Limit(int(query.Limit.Int64))
	}
	tx = tx.Offset(query.Offset)
	return tx
}

// buildFromRequestQueryFilter dynamically build a query filter from the given requests.QueryFilter
func buildFromRequestQueryFilter(tx *gorm.DB, filter requests.QueryFilter) *gorm.DB {
	switch filter.Operator {
	// Equals query
	case enums.QueryFilterOperatorEquals:
		tx = tx.Where(filter.Column+" = ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrEquals:
		tx = tx.Or(filter.Column+" = ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorNotEquals:
		tx = tx.Not(filter.Column+" = ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrNotEquals:
		tx = tx.Or("NOT "+filter.Column+" = ?", filter.Values[0])
		break
	// Contains query
	case enums.QueryFilterOperatorContains:
		tx = tx.Where(filter.Column+" LIKE ?", fmt.Sprintf("%s%v%s", "%", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorOrContains:
		tx = tx.Or(filter.Column+" LIKE ?", fmt.Sprintf("%s%v%s", "%", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorNotContains:
		tx = tx.Not(filter.Column+" LIKE ?", fmt.Sprintf("%s%v%s", "%", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorOrNotContains:
		tx = tx.Or("NOT "+filter.Column+" LIKE ?", fmt.Sprintf("%s%v%s", "%", filter.Values[0], "%"))
		break
	// Starts with query
	case enums.QueryFilterOperatorStartsWith:
		tx = tx.Where(filter.Column+" LIKE ?", fmt.Sprintf("%v%s", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorOrStartsWith:
		tx = tx.Or(filter.Column+" LIKE ?", fmt.Sprintf("%v%s", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorNotStartsWith:
		tx = tx.Not(filter.Column+" LIKE ?", fmt.Sprintf("%v%s", filter.Values[0], "%"))
		break
	case enums.QueryFilterOperatorOrNotStartsWith:
		tx = tx.Or("NOT "+filter.Column+" LIKE ?", fmt.Sprintf("%v%s", filter.Values[0], "%"))
		break
	// Ends with query
	case enums.QueryFilterOperatorEndsWith:
		tx = tx.Where(filter.Column+" LIKE ?", fmt.Sprintf("%s%v", "%", filter.Values[0]))
		break
	case enums.QueryFilterOperatorOrEndsWith:
		tx = tx.Or(filter.Column+" LIKE ?", fmt.Sprintf("%s%v", "%", filter.Values[0]))
		break
	case enums.QueryFilterOperatorNotEndsWith:
		tx = tx.Not(filter.Column+" LIKE ?", fmt.Sprintf("%s%v", "%", filter.Values[0]))
		break
	case enums.QueryFilterOperatorOrNotEndsWith:
		tx = tx.Or("NOT "+filter.Column+" LIKE ?", fmt.Sprintf("%s%v", "%", filter.Values[0]))
		break
	// Greater than query
	case enums.QueryFilterOperatorGreaterThan:
		tx = tx.Where(filter.Column+" > ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrGreaterThan:
		tx = tx.Or(filter.Column+" > ?", filter.Values[0])
		break
	// Greater than equal query
	case enums.QueryFilterOperatorGreaterThanEqual:
		tx = tx.Where(filter.Column+" >= ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrGreaterThanEqual:
		tx = tx.Or(filter.Column+" >= ?", filter.Values[0])
		break
	// Lesser than query
	case enums.QueryFilterOperatorLesserThan:
		tx = tx.Where(filter.Column+" < ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrLesserThan:
		tx = tx.Or(filter.Column+" < ?", filter.Values[0])
		break
	// Lesser than equal query
	case enums.QueryFilterOperatorLesserThanEqual:
		tx = tx.Where(filter.Column+" <= ?", filter.Values[0])
		break
	case enums.QueryFilterOperatorOrLesserThanEqual:
		tx = tx.Or(filter.Column+" <= ?", filter.Values[0])
		break
	// In array query
	case enums.QueryFilterOperatorIn:
		tx = tx.Where(filter.Column+" IN ?", filter.Values...)
		break
	case enums.QueryFilterOperatorOrIn:
		tx = tx.Or(filter.Column+" IN ?", filter.Values...)
		break
	case enums.QueryFilterOperatorNotIn:
		tx = tx.Not(filter.Column+" IN ?", filter.Values...)
		break
	case enums.QueryFilterOperatorOrNotIn:
		tx = tx.Or("NOT "+filter.Column+" IN ?", filter.Values...)
		break
	// Between query
	case enums.QueryFilterOperatorBetween:
		tx = tx.Where(filter.Column+" BETWEEN ? AND ?", filter.Values[0:1]...)
		break
	case enums.QueryFilterOperatorOrBetween:
		tx = tx.Or(filter.Column+" BETWEEN ? AND ?", filter.Values[0:1]...)
		break
	case enums.QueryFilterOperatorNotBetween:
		tx = tx.Not(filter.Column+" BETWEEN ? AND ?", filter.Values[0:1]...)
		break
	case enums.QueryFilterOperatorOrNotBetween:
		tx = tx.Or("NOT "+filter.Column+" BETWEEN ? AND ?", filter.Values[0:1]...)
		break
	}
	return tx
}
