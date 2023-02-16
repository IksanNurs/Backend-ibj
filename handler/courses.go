package handler

import (
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/courses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CoursesHandler struct {
	coursesService courses.Service
}

func NewCoursesHandler(coursesService courses.Service) *CoursesHandler {
	return &CoursesHandler{coursesService}
}

func (h *CoursesHandler) Create(c *gin.Context) {
	var input courses.InputCourses
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan courses", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.coursesService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal input data courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := courses.FormatCourses(data)
	response := helper.APIResponse("Berhasil menambah data courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CoursesHandler) Read(c *gin.Context) {
	data, err := h.coursesService.Read()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal menampilkan data courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	formatter := courses.FormatsCourses(data)
	response := helper.APIResponse("Berhasil menampilkan data courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CoursesHandler) Update(c *gin.Context) {
	var input courses.UpdateCourses
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan courses", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.coursesService.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal update data courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := courses.FormatCourses(data)
	response := helper.APIResponse("Berhasil update data courses", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CoursesHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.coursesService.Delete(ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal delete data courses", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	response := helper.APIResponse("Berhasil delete data courses", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}
