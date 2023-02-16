package admin

type RegisterInputAdmin struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInputAdmin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
