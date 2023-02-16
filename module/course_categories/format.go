package course_categories

type FormatKey struct {
	Course_categories interface{} `json:"course_categories"`
}

type SubFormatKeyValue struct {
	ID    int    `json:"id" sql:"unique"`
	Name  string `json:"name"`
}


func FormatCourseCategories(course_categories CourseCategories) FormatKey {
	
	formatter:=FormatKey{
	Course_categories: SubFormatKeyValue{
		ID: int(course_categories.ID.Int64),
		Name: course_categories.Name.String,
	},
	}

	return formatter
}

func FormatsCourseCategories(course_categories []CourseCategories) FormatKey {
	var formatters []SubFormatKeyValue

    for _, course_category:= range course_categories{
		formatter:=SubFormatKeyValue{
			ID: int(course_category.ID.Int64),
			Name: course_category.Name.String,
		}
		formatters=append(formatters, formatter)     
	}
    
   form:=FormatKey{
	Course_categories: formatters,
   }
	return form
}