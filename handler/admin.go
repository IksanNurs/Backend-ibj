package handler

import (
	"backend_iksan_nursalim/auth"
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/admin"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService admin.Service
	authService  auth.Service
}

func NewAdminHandler(adminService admin.Service, authService auth.Service) *AdminHandler {
	return &AdminHandler{adminService, authService}
}

func (h *AdminHandler) Register(c *gin.Context) {
	var input admin.RegisterInputAdmin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("gagal validasi inputan register", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newAdmin, err := h.adminService.Register(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("gagal input data register", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println(err)
		return
	}

	token, err := h.authService.GenerateToken(int(newAdmin.ID.Int64))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("gagal generate token", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := admin.FormatAdmin(newAdmin, token)
	response := helper.APIResponse("Berhasil register", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *AdminHandler) Login(c *gin.Context) {
	var input admin.LoginInputAdmin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("gagal validasi inputan login", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.adminService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("gagal comparasi data login", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(int(loggedinUser.ID.Int64))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("gagal memperbarui token", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := admin.FormatAdmin(loggedinUser, token)
	response := helper.APIResponse("Berhasil login", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}
