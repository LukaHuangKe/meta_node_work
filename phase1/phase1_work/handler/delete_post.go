package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	var req pb.DeletePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleDeletePost(c, &req))
}

func handleDeletePost(ctx *gin.Context, req *pb.DeletePostReq) (resp *pb.DeletePostResp) {
	resp = &pb.DeletePostResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	postId, err := strconv.ParseInt(req.GetPostId(), 10, 64)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	if err := mysql_logic.DeletePost(ctx, postId, userID); err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	return resp
}
