package handler

import (
	"phase1/phase1_work/pb"

	"github.com/gin-gonic/gin"
)

func ListCommentByPostId(c *gin.Context) {
	var req pb.ListCommentByPostIdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	c.JSON(200, &pb.ListCommentByPostIdResp{})
}
