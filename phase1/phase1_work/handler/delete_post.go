package handler

import (
	"phase1/phase1_work/pb"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	var req pb.DeletePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	c.JSON(200, &pb.DeletePostResp{})
}
