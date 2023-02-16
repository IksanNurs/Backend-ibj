package courses

import (
	"backend_iksan_nursalim/module/course_categories"
	"database/sql"
)

type Courses struct {
	ID                  sql.NullInt64
	Title               sql.NullString
	Courses_category_id sql.NullInt64
	Course_categories    course_categories.CourseCategories
}
