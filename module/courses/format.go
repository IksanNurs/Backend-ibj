package courses

import "backend_iksan_nursalim/module/course_categories"

type FormatKey struct {
	Courses interface{} `json:"courses"`
}

type SubFormatKeyValue struct {
	ID                  int         `json:"id" sql:"unique"`
	Title               string      `json:"title"`
	Courses_category_id int         `json:"courses_category_id"`
	Course_categories   interface{} `json:"course_categories"`
}

func FormatCourses(courses Courses) FormatKey {

	formatter := FormatKey{
		Courses: SubFormatKeyValue{
			ID:                  int(courses.ID.Int64),
			Title:               courses.Title.String,
			Courses_category_id: int(courses.Courses_category_id.Int64),
			Course_categories: course_categories.SubFormatKeyValue{
				ID: int(courses.Course_categories.ID.Int64),
				Name: courses.Course_categories.Name.String,
			},
		},
	}

	return formatter
}

func FormatsCourses(courses []Courses) FormatKey {
	var formatters []SubFormatKeyValue

	for _, course := range courses {
		formatter := SubFormatKeyValue{
			ID:                  int(course.ID.Int64),
			Title:               course.Title.String,
			Courses_category_id: int(course.Courses_category_id.Int64),
			Course_categories:   course_categories.SubFormatKeyValue{
				ID: int(course.Course_categories.ID.Int64),
				Name: course.Course_categories.Name.String,
			},
		}
		formatters = append(formatters, formatter)
	}

	form := FormatKey{
		Courses: formatters,
	}
	return form
}
