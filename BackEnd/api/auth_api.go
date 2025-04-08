package api

import (
	"BackEnd/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RegisterRequest struct {
	Address string `json:"address"`
	Role    string `json:"role"`
}

type LoginRequest struct {
	Address string `json:"address"`
}

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Role string `json:"role"`
	} `json:"data"`
}

// 处理登录请求
func LoginWithMetaMask(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
		return
	}

	role, err := service.GetUserRole(req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户角色失败"})
		return
	}

	if role == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "用户未找到"})
		return
	}

	res := LoginResponse{
		Code: 200,
		Msg:  "登录成功",
	}
	res.Data.Role = role
	c.JSON(http.StatusOK, res)
}

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
			return
		}

		err := service.RegisterUser(db, req.Address, req.Role)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	}
}
