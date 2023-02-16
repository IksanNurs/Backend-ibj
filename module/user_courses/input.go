package user_courses

type InputUserCourses struct {
	Users_id  int `json:"users_id" binding:"required"`
	Course_id int `json:"course_id" binding:"required"`
}

type UpdateUserCourses struct {
	ID  int `json:"ID" binding:"required"`
	Users_id  int `json:"users_id" binding:"required"`
	Course_id int `json:"course_id" binding:"required"`
}