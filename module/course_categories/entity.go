package course_categories

import "database/sql"

type CourseCategories struct {
	ID sql.NullInt64
	Name sql.NullString
}