package handler

import (
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/user_courses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCoursesHandler struct {
	userCoursesService user_courses.Service
}

func NewUserCoursesHandler(userCoursesService user_courses.Service) *UserCoursesHandler {
	return &UserCoursesHandler{userCoursesService}
}

func (h *UserCoursesHandler) Create(c *gin.Context) {
	var input user_courses.InputUserCourses
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan user_courses", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.userCoursesService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal input data user_courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := user_courses.FormatUserCourses(data)
	response := helper.APIResponse("Berhasil menambah data user_courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UserCoursesHandler) Read(c *gin.Context) {
	data, err := h.userCoursesService.Read()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal menampilkan data user_courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	formatter := user_courses.FormatsUserCourses(data)
	response := helper.APIResponse("Berhasil menampilkan data user_courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UserCoursesHandler) Update(c *gin.Context) {
	var input user_courses.UpdateUserCourses
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan user_courses", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.userCoursesService.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal update data user_courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := user_courses.FormatUserCourses(data)
	response := helper.APIResponse("Berhasil update data user_courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UserCoursesHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.userCoursesService.Delete(ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal delete data user_courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	response := helper.APIResponse("berhasil delete data user_courses", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}
