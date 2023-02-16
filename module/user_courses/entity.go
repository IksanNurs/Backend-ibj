package user_courses

import (
	"backend_iksan_nursalim/module/course_categories"
	"backend_iksan_nursalim/module/courses"
	"backend_iksan_nursalim/module/users"
	"database/sql"
)

type UserCourses struct {
	ID              sql.NullInt64
	Users_id        sql.NullInt64
	Course_id       sql.NullInt64
	Users           users.Users
	Courses         courses.Courses
	Course_category course_categories.CourseCategories
}
