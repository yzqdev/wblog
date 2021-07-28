package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gookit/color"
	"net/http"
	"wblog/models"
)

func JwtHandler() gin.HandlerFunc {
	//@Security ApiKeyAuth
	return func(context *gin.Context) {
		result := models.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证，重新登录",
			Data:    nil,
		}
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}

		// 校验token
		adminUser, err := parseToken(auth)
		context.Set("user", adminUser)

		if err != nil {
			context.Abort()
			result.Message = err.Error()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		} else {
			println("token 正确")
		}
		context.Next()
	}
}
func parseToken(yourToken string) (models.User, error) {
	claims := models.NewJwtClaims{}
	_, err := jwt.ParseWithClaims(yourToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return controller.SecretKey, nil

	})
	if err != nil {
		color.Danger.Println("token值为空")

	}
	color.Danger.Println(claims.AdminUser, "编译token")
	return *claims.AdminUser, err

}
