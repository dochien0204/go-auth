package handlersLogin

import (
	"fmt"
	"jwt-project/configs"
	"jwt-project/controllers/login"
	"jwt-project/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service login.Service
}

func NewHandlerLogin(service login.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input login.InputLogin

	ctx.ShouldBindJSON(&input)

	resultLogin, errorLogin := h.service.LoginService(&input)
	fmt.Println("Error", errorLogin)

	if errorLogin == "Invalid account" {
		utils.ValidatorErrorResponse(ctx, "Login Failed", http.StatusBadRequest, http.MethodPost, errorLogin)
		return
	}

	accessTokenData := map[string]interface{}{"username": resultLogin.Username}
	timeExpired := 10 * time.Minute
	accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", timeExpired)

	//redis to save token is used
	redisClient := configs.RedisDatabase()

	//save to redis
	redisClient.LPush(configs.Ctx, "accessToken", accessToken)

	//check range of list in redis
	lenList := redisClient.LLen(configs.Ctx, "accessToken")
	lenRange := redisClient.LRange(configs.Ctx, "accessToken", 0, lenList.Val())

	fmt.Println("List Redis:", lenRange)

	if errToken != nil {
		utils.APIResponse(ctx, "Generate token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, accessToken)
}
