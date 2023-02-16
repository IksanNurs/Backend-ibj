package handler

import (
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/users"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	usersService users.Service
}

func NewUsersHandler(usersService users.Service) *UsersHandler {
	return &UsersHandler{usersService}
}

func (h *UsersHandler) Create(c *gin.Context) {
	var input users.InputUsers
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan users", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.usersService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal input data users", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := users.FormatUsers(data)
	response := helper.APIResponse("Berhasil menambahkan data users", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Read(c *gin.Context) {
	data, err := h.usersService.Read()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal menampilkan data users", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	formatter := users.FormatsUsers(data)
	response := helper.APIResponse("Berhasil menampilkan data users", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Update(c *gin.Context) {
	var input users.UpdateUsers
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal validasi inputan users", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.usersService.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal update data users", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	formatter := users.FormatUsers(data)
	response := helper.APIResponse("Berhasil update data users", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.usersService.Delete(ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Gagal delete data users", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}
	response := helper.APIResponse("Berhasil delete data users", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}
