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

func CreatePost(c *gin.Context) {
	var req pb.CreatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleCreatePost(c, &req))
}

func handleCreatePost(ctx *gin.Context, req *pb.CreatePostReq) (resp *pb.CreatePostResp) {
	resp = &pb.CreatePostResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	if len(req.GetPost().GetTitle()) == 0 || len(req.GetPost().GetContent()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = utils.FailCodeMsg
	}

	post := &mysql_model.Posts{
		Title:     req.GetPost().GetTitle(),
		Content:   req.GetPost().GetContent(),
		UserId:    userID,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
		DeletedAt: 0,
	}

	postId, err := mysql_logic.CreatePost(ctx, post)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}
	resp.PostId = strconv.FormatInt(postId, 10)
	return resp
}
