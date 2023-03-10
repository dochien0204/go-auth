package logoutHandler

import (
	"fmt"
	"jwt-project/configs"
	"jwt-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(ctx *gin.Context) {
	// token := ctx.Param("token")

	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	accessToken = strings.Trim(accessToken, " ")

	fmt.Println("AccessToken:", accessToken)
	redisClient := configs.RedisDatabase()

	redisClient.LRem(configs.Ctx, "accessToken", -1, accessToken)
	lenList := redisClient.LLen(configs.Ctx, "accessToken")
	lenRange := redisClient.LRange(configs.Ctx, "accessToken", 0, lenList.Val())
	fmt.Println(lenRange)

	utils.APIResponse(ctx, "Logout succesfully", http.StatusOK, http.MethodPost, nil)
}
