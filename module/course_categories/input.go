package course_categories

type InputCourseCategories struct {
	Nama string `json:"name" binding:"required"`
}

type UpdateCourseCategories struct {
	ID   int    `json:"id" binding:"required"`
	Nama string `json:"name" binding:"required"`
}
