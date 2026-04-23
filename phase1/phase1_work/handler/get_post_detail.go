package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPostDetail(c *gin.Context) {
	var req pb.GetPostDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleGetPostDetail(c, &req))
}

func handleGetPostDetail(ctx *gin.Context, req *pb.GetPostDetailReq) (resp *pb.GetPostDetailResp) {
	resp = &pb.GetPostDetailResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	postId, err := strconv.ParseInt(req.GetPostId(), 10, 64)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	post, err := mysql_logic.GetPostDetail(ctx.Request.Context(), postId)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	// 获取作者信息
	author, err := mysql_logic.GetUserById(ctx.Request.Context(), post.UserId)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	resp.Post = &pb.Post{
		PostId:     strconv.FormatInt(post.Id, 10),
		Title:      post.Title,
		Content:    post.Content,
		CreateTime: post.CreatedAt,
		UpdateTime: post.UpdatedAt,
		Author: &pb.User{
			UserId:   strconv.FormatInt(author.Id, 10),
			Username: author.Username,
			Email:    author.Email,
		},
	}

	return resp
}
