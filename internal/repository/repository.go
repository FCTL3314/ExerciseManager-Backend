package repository

import "gorm.io/gorm"

func applyPreloadsForGORMQuery(query *gorm.DB, preloads []string) *gorm.DB {
	for _, preload := range preloads {
		query.Preload(preload)
	}
	return query
}
