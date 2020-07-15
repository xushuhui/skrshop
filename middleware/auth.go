package middleware

import (
	"github.com/gin-gonic/gin"
	cd "skrshop-api/code"
	"skrshop-api/core"
	"skrshop-api/utils"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = 0
		token := c.Query("token")
		if token == "" {
			code = cd.InvalidParams
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = cd.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = cd.TimeoutAuthToken
			}
		}

		if code != 0 {
			core.FailResp(c, code)
			return
		}

		c.Next()
	}

}

func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	m := make(map[string]interface{})
	m["token"] = token
	m["expire"] = expire
	core.SetData(c, m)
}
