package auth

import (
	"baby-daily-api/configs"
	"baby-daily-api/internal/response"
	"baby-daily-api/internal/service/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleWare 登录认证中间件
func AuthMiddleWare(jwtService *jwt.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c)
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, "缺少token")
			c.Abort()
			return
		}

		user, err := jwtService.ParseToken(token)
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, "非法token")
			c.Abort()
			return
		}

		c.Set(configs.UserContextKey, user)
		c.Next()

	}
}

// 从token中获取认证信息
func getTokenFromHeader(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("无效token")
	}
	token := strings.Fields(auth)
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return "", errors.New("无效token")
	}
	return token[1], nil
}
