package handler

import (
	"phase1/phase1_work/pb"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// 你的业务逻辑
	c.JSON(200, &pb.LoginResp{})
}
