package admin_service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/middleware"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

//todo:并发锁 sync.map
//var AdminloginMap = map[string]string{} /*记录登录用户信息 */
var AdminloginMap sync.Map /*记录登录用户信息 */

//检查是否和当前存储的token是否一致
func CheckCurrentToken(userID string, token string) cerror.Cerror {
	str, ok := AdminloginMap.Load(userID)
	if !ok { //未有登录记录
		return cerror.ErrorUserNotLogin
	}

	if str != token { //token已经更新
		return cerror.ErrorTokenAuthFailed

	}
	//已经登录且token匹配
	return cerror.ErrorUserAuthSucc
}

//检查登录是否过期
func CheckExpired(tokenStr string) int {
	token, err := jwt.ParseWithClaims(tokenStr, &middleware.CustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.SecretKey), nil
	})

	if claims, ok := token.Claims.(*middleware.CustomClaimsAdmin); ok && token.Valid {
		//fmt.Println(claims)
		fmt.Printf("%+v", claims)
		fmt.Println(claims.ExpiresAt, time.Now().Unix())

		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			fmt.Println("未过期")
			return common.NotExpired
		} else {
			fmt.Println("已过期。。。")
			return common.TokenExpired
		}

	} else {
		fmt.Println(err)
	}

	return common.NotExpired
}

//检查用户名和密码
func UserLogin(ctx *gin.Context, name string, password string) (string, cerror.Cerror) {
	var res cerror.Cerror = cerror.ErrorLoginSucc
	err := dao.CheckAdmin(ctx, name, password)

	var newTokenStr string
	if err.Code() == cerror.ErrorUserAuthSucc.Code() { //用户名，密码正确

		newTokenStr, _ = middleware.GenerateTokenAdmin(name)

		//检查是否已经登录，
		if oldTokenStr, ok := AdminloginMap.Load(name); ok {
			logger.Debugf("%s", oldTokenStr)
			//如果已经登录，则检查token是否过期
			if common.NotExpired == CheckExpired(oldTokenStr.(string)) {
				res = cerror.ErrorLoginAgain
			}
		}

		AdminloginMap.Store(name, newTokenStr) //无论是不是已经登录，都更新；如果已经登录，则之前的登录无效

	} else {
		return "", cerror.ErrorLoginFailed
	}

	return newTokenStr, res
}

func Login(ctx *gin.Context, userName string, password string) (string, cerror.Cerror) {

	tokenStr, res := UserLogin(ctx, userName, password)
	logger.Debugf("tokenStr:%s,%d", tokenStr, res)

	return tokenStr, res
}

func Logout(ctx *gin.Context, tokenStr string) int {

	AdminloginMap.Range(func(k, v interface{}) bool {

		// 这里可以根据需要对每个键值对进行操作
		fmt.Printf("Key: %s, Value: %s\n", k.(string), v.(string))
		return true
	})

	AdminloginMap.Range(func(k, v interface{}) bool {
		if v.(string) == tokenStr {
			AdminloginMap.Delete(k)
		}
		return true
	})

	return common.Success

}
