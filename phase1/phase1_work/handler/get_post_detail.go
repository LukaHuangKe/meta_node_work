package handler

import (
	"phase1/phase1_work/pb"

	"github.com/gin-gonic/gin"
)

func GetPostDetail(c *gin.Context) {
	var req pb.GetPostDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, &pb.GetPostDetailResp{})
}
