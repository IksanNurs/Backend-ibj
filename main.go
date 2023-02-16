package main

import (
	"backend_iksan_nursalim/auth"
	"backend_iksan_nursalim/database"
	"backend_iksan_nursalim/handler"
	"backend_iksan_nursalim/middleware"
	"backend_iksan_nursalim/module/admin"
	"backend_iksan_nursalim/module/course_categories"
	"backend_iksan_nursalim/module/courses"
	"backend_iksan_nursalim/module/user_courses"
	"backend_iksan_nursalim/module/users"
	"backend_iksan_nursalim/utils"
	_ "embed"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed .env
var env string

func main() {
	utils.LoadEnv(env)
	db := database.Database()
	defer db.Close()

	adminRepository := admin.NewRepository(db)
	adminService := admin.NewService(adminRepository)
	authService := auth.NewService()
	adminHandler := handler.NewAdminHandler(adminService, authService)

	usersRepository := users.NewRepository(db)
	usersService := users.NewService(usersRepository)
	usersHandler := handler.NewUsersHandler(usersService)

	userCoursesRepository := user_courses.NewRepository(db)
	usersCoursesService := user_courses.NewService(userCoursesRepository)
	usersCoursesHandler := handler.NewUserCoursesHandler(usersCoursesService)

	coursesRepository := courses.NewRepository(db)
	coursesService := courses.NewService(coursesRepository)
	coursesHandler := handler.NewCoursesHandler(coursesService)

	courseCategoriesRepository := course_categories.NewRepository(db)
	courseCategoriesService := course_categories.NewService(courseCategoriesRepository)
	courseCategoriesHandler := handler.NewCourseCategoriesHandler(courseCategoriesService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := router.Group("/api")

	//====Admin====
	api.POST("/register", adminHandler.Register)
	api.POST("/login", adminHandler.Login)

	//====Users====
	api.POST("/users", middleware.AuthMiddleware(authService, adminService), usersHandler.Create)
	api.GET("/users", middleware.AuthMiddleware(authService, adminService), usersHandler.Read)
	api.PUT("/users", middleware.AuthMiddleware(authService, adminService), usersHandler.Update)
	api.DELETE("/users/:id", middleware.AuthMiddleware(authService, adminService), usersHandler.Delete)

	//====User_Courses====
	api.POST("/user-courses", middleware.AuthMiddleware(authService, adminService), usersCoursesHandler.Create)
	api.GET("/user-courses", middleware.AuthMiddleware(authService, adminService), usersCoursesHandler.Read)
	api.PUT("/user-courses", middleware.AuthMiddleware(authService, adminService), usersCoursesHandler.Update)
	api.DELETE("/user-courses/:id", middleware.AuthMiddleware(authService, adminService), usersCoursesHandler.Delete)

	//====Courses====
	api.POST("/courses", middleware.AuthMiddleware(authService, adminService), coursesHandler.Create)
	api.GET("/courses", middleware.AuthMiddleware(authService, adminService), coursesHandler.Read)
	api.PUT("/courses", middleware.AuthMiddleware(authService, adminService), coursesHandler.Update)
	api.DELETE("/courses/:id", middleware.AuthMiddleware(authService, adminService), coursesHandler.Delete)

	//====Courses_Categories====
	api.POST("/course-categories", middleware.AuthMiddleware(authService, adminService), courseCategoriesHandler.Create)
	api.GET("/course-categories", middleware.AuthMiddleware(authService, adminService), courseCategoriesHandler.Read)
	api.PUT("/course-categories", middleware.AuthMiddleware(authService, adminService), courseCategoriesHandler.Update)
	api.DELETE("/course-categories/:id", middleware.AuthMiddleware(authService, adminService), courseCategoriesHandler.Delete)

	router.Run(":" + os.Getenv("APP_PORT"))
}
