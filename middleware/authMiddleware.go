package middleware

import (
	"backend_iksan_nursalim/auth"
	"backend_iksan_nursalim/helper"
	"backend_iksan_nursalim/module/admin"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service, adminService admin.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error bearer")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Unauthorized1", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized2", http.StatusUnauthorized, "error token invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		adminID := int(claim["user_id"].(float64))
		admin, err := adminService.FindByID(adminID)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Unauthorized3", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentAdmin", admin)

	}
}
