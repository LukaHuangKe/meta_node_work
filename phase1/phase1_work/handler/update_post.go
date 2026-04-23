package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	var req pb.UpdatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleUpdatePost(c, &req))
}

func handleUpdatePost(ctx *gin.Context, req *pb.UpdatePostReq) (resp *pb.UpdatePostResp) {
	resp = &pb.UpdatePostResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	if req.Post == nil || len(req.Post.GetTitle()) == 0 || len(req.Post.GetContent()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = "title and content cannot be empty"
		return resp
	}

	postId, err := strconv.ParseInt(req.GetPost().GetPostId(), 10, 64)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	if err := mysql_logic.UpdatePost(ctx.Request.Context(), postId, userID, req.GetPost().GetTitle(), req.GetPost().GetContent()); err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	return resp
}
