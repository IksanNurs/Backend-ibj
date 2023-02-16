package user_courses

import (
	"backend_iksan_nursalim/module/course_categories"
	"backend_iksan_nursalim/module/courses"
	"backend_iksan_nursalim/module/users"
)

type FormatKey struct {
	User_courses interface{} `json:"user_courses"`
}

type SubFormatKeyValue struct {
	ID        int         `json:"id" sql:"unique"`
	Users_id  int         `json:"users_id"`
	Course_id int         `json:"course_id"`
	Users     interface{} `json:"users"`
	Courses   interface{} `json:"courses"`
}

func FormatUserCourses(user_courses UserCourses) FormatKey {

	formatter := FormatKey{
		User_courses: SubFormatKeyValue{
			ID:        int(user_courses.ID.Int64),
			Users_id:  int(user_courses.Users_id.Int64),
			Course_id: int(user_courses.Course_id.Int64),
			Users: users.SubFormatKeyValue{
				ID:    int(user_courses.Users.ID.Int64),
				Name:  user_courses.Users.Name.String,
				Email: user_courses.Users.Email.String,
			},
			Courses: courses.SubFormatKeyValue{
				ID: int(user_courses.Courses.ID.Int64),
				Title: user_courses.Courses.Title.String,
				Courses_category_id: int(user_courses.Courses.Courses_category_id.Int64),
				Course_categories: course_categories.SubFormatKeyValue{
					ID: int(user_courses.Course_category.ID.Int64),
					Name: user_courses.Course_category.Name.String,
				},
			},
		},
	}

	return formatter
}

func FormatsUserCourses(user_courses []UserCourses) FormatKey {
	var formatters []SubFormatKeyValue

	for _, user_course := range user_courses {
		formatter := SubFormatKeyValue{
			ID:        int(user_course.ID.Int64),
			Users_id:  int(user_course.Users_id.Int64),
			Course_id: int(user_course.Course_id.Int64),
			Users: users.SubFormatKeyValue{
				ID:    int(user_course.Users.ID.Int64),
				Name:  user_course.Users.Name.String,
				Email: user_course.Users.Email.String,
			},
			Courses: courses.SubFormatKeyValue{
				ID: int(user_course.Courses.ID.Int64),
				Title: user_course.Courses.Title.String,
				Courses_category_id: int(user_course.Courses.Courses_category_id.Int64),
				Course_categories: course_categories.SubFormatKeyValue{
					ID: int(user_course.Course_category.ID.Int64),
					Name: user_course.Course_category.Name.String,
				},
			},
		}
		formatters = append(formatters, formatter)
	}

	form := FormatKey{
		User_courses: formatters,
	}
	return form
}
