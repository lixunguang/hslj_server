package middleware

import (
<<<<<<< HEAD
	"hslj/pkg/cerror"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"fmt"
	"github.com/golang-jwt/jwt"
	"hslj/pkg/cerror"
	"time"
)

var SecretKey = []byte("YundaoEdu")

type CustomClaims struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.StandardClaims
}

type CustomClaimsAdmin struct {
	UserName string `json:"user_name"`
	//Exp      int64  `json:"exp"`
	jwt.StandardClaims
}

//后台管理
func GenerateTokenAdmin(username string) (string, error) {

	//  Create the Claims
	claims := CustomClaimsAdmin{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 7 * 60).Unix(),
			Issuer:    fmt.Sprintf("%d", time.Now().Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

//根据token
func ParseAdminToken(tokenString string) (string, cerror.Cerror) {
	fmt.Println("token is=", tokenString)

	var claims CustomClaimsAdmin
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	if token == nil {
		return "", cerror.ErrorTokenFormat
	}

	if claims, ok := token.Claims.(*CustomClaimsAdmin); ok {
		isExpires := claims.VerifyExpiresAt(time.Now().Unix(), true)
		if isExpires {
			fmt.Println("Token not expired.")
			return claims.UserName, nil
		} else {
			return "", cerror.ErrorTokenExpire
		}
	} else {
		if ve, ok := err.(*jwt.ValidationError); ok { //官方写法
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
				return "", cerror.ErrorTokenFormat
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期或未启用")

				return "", cerror.ErrorTokenExpire
			} else {
				fmt.Println("无法处理这个token", err)
				return "", cerror.ErrorTokenFormat
			}
		}
	}

	return "", cerror.ErrorTokenExpire
}
