package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/mysql/mysql_model"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var req pb.CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleCreateComment(c, &req))
}

func handleCreateComment(ctx *gin.Context, req *pb.CreateCommentReq) (resp *pb.CreateCommentResp) {
	resp = &pb.CreateCommentResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	if len(req.GetComment().GetContent()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = "content cannot be empty"
		return resp
	}

	postId, err := strconv.ParseInt(req.GetComment().GetPostId(), 10, 64)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = "invalid post id"
		return resp
	}

	comment := &mysql_model.Comments{
		Content:   req.GetComment().GetContent(),
		UserId:    userID,
		PostId:    postId,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
		DeletedAt: 0,
	}

	commentId, err := mysql_logic.CreateComment(ctx.Request.Context(), comment)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	resp.CommentId = strconv.FormatInt(commentId, 10)
	return resp
}
