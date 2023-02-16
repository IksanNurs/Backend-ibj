package handler

import (
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/course_categories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseCategoriesHandler struct {
	courseCategoriesService course_categories.Service
}

func NewCourseCategoriesHandler(courseCategoriesService course_categories.Service) *CourseCategoriesHandler {
	return &CourseCategoriesHandler{courseCategoriesService}
}

func (h *CourseCategoriesHandler) Create(c *gin.Context) {
	var input course_categories.InputCourseCategories
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan course_categories", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.courseCategoriesService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal input data course_categories", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := course_categories.FormatCourseCategories(data)
	response := helper.APIResponse("Berhasil menambahkan data course_categories", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CourseCategoriesHandler) Read(c *gin.Context) {
	data, err := h.courseCategoriesService.Read()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal menampilkan data course_categories", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	formatter := course_categories.FormatsCourseCategories(data)
	response := helper.APIResponse("berhasil menampilkan data course_categories", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CourseCategoriesHandler) Update(c *gin.Context) {
	var input course_categories.UpdateCourseCategories
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan course_categories", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.courseCategoriesService.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal update data course_categories", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := course_categories.FormatCourseCategories(data)
	response := helper.APIResponse("Berhasil update data course_categories", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CourseCategoriesHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.courseCategoriesService.Delete(ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal delete data course_categories", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	response := helper.APIResponse("Berhasil delete data course_categories", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}
