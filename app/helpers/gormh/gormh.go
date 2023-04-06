// GORM helper package`
package gormh

import (
	"fmt"
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"gorm.io/gorm"
)

// BuildFromRequestQuery builds a gorm query from the given requests.Query
func BuildFromRequestQuery(db *gorm.DB, query requests.Query) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	scopes := [](func(*gorm.DB) *gorm.DB){}
	for _, filter := range query.Filters {
		syncronizer.WG().Add(1)
		go func(filter requests.QueryFilter) {
			defer syncronizer.WG().Done()
			var scope func(*gorm.DB) *gorm.DB

			switch filter.Operator {
			case "==", "=":
				scope = func(*gorm.DB) *gorm.DB {
					db = db.Where(filter.Column+" = ?", filter.Values[0])
					// Loop start from index one not index zero
					for i := 1; i < len(filter.Values); i++ {
						db.Or(filter.Column+" = ?", filter.Values[i])
					}
					return db
				}
				break
			case "!=", "!":
				scope = func(*gorm.DB) *gorm.DB {
					db = db.Not(filter.Column+" = ?", filter.Values[0])
					// Loop start from index one not index zero
					for i := 1; i < len(filter.Values); i++ {
						db.Or(filter.Column+" = ?", filter.Values[i])
					}
					return db
				}
				break
			case ">", ">=", "<", "<=":
				scope = func(*gorm.DB) *gorm.DB {
					return db.Where(filter.Column+" "+filter.Operator+" ?", filter.Values[0])
				}
				break
			case "--", "-":
				scope = func(*gorm.DB) *gorm.DB {
					return db.Where(filter.Column+" BETWEEN ? AND ?", filter.Values[0], filter.Values[1])
				}
				break
			case ".%", "LIKE%":
				scope = func(*gorm.DB) *gorm.DB {
					return db.Where(
						filter.Column+" "+filter.Operator+" ?",
						fmt.Sprintf("%v", filter.Values[0])+"%",
					)
				}
				break
			case "%.", "%LIKE":
				scope = func(*gorm.DB) *gorm.DB {
					return db.Where(
						filter.Column+" "+filter.Operator+" ?",
						"%"+fmt.Sprintf("%v", filter.Values[0]),
					)
				}
				break
			case "%%", "%.%", "%LIKE%":
				scope = func(*gorm.DB) *gorm.DB {
					return db.Where(
						filter.Column+" "+filter.Operator+" ?",
						"%"+fmt.Sprintf("%v", filter.Values[0])+"%",
					)
				}
				break
			}
			syncronizer.M().Lock()
			scopes = append(scopes, scope)
			syncronizer.M().Unlock()
		}(filter)
	}
	for column, orderingType := range query.OrderBys {
		syncronizer.WG().Add(1)
		go func(column, orderingType string) {
			defer syncronizer.WG().Done()
			syncronizer.M().Lock()
			db.Order(column + " " + strings.ToLower(orderingType))
			syncronizer.M().Unlock()
		}(column, orderingType)
	}
	for _, with := range query.Withs {
		syncronizer.WG().Add(1)
		go func(with string) {
			defer syncronizer.WG().Done()
			syncronizer.M().Lock()
			db.Preload(with)
			syncronizer.M().Unlock()
		}(with)
	}
	syncronizer.WG().Wait()
	if len(scopes) != 0 {
		db = db.Scopes(scopes...)
	}
	db = db.Offset(query.Offset).Limit(query.Limit)
}
