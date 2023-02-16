package courses

type InputCourses struct {
	Title              string `json:"title" binding:"required"`
	Course_category_id int `json:"course_category_id" binding:"required"`
}

type UpdateCourses struct {
	ID                 int    `json:"id" binding:"required"`
	Title              string `json:"title" binding:"required"`
	Course_category_id int `json:"course_category_id" binding:"required"`
}
