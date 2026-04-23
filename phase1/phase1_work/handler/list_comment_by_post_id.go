package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/mysql/mysql_model"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListCommentByPostId(c *gin.Context) {
	var req pb.ListCommentByPostIdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleListCommentByPostId(c, &req))
}

func handleListCommentByPostId(ctx *gin.Context, req *pb.ListCommentByPostIdReq) (resp *pb.ListCommentByPostIdResp) {
	resp = &pb.ListCommentByPostIdResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	// postId 必传校验
	if len(req.GetPostId()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = "post id is required"
		return resp
	}

	// 解析 postId
	postId, err := strconv.ParseInt(req.GetPostId(), 10, 64)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	// 分页参数
	pageNo := int32(1)
	pageSize := int32(10)
	if req.Pagination != nil {
		if req.Pagination.PageNo > 0 {
			pageNo = req.Pagination.PageNo
		}
		if req.Pagination.PageSize > 0 && req.Pagination.PageSize <= 100 {
			pageSize = req.Pagination.PageSize
		}
	}

	// 查询评论列表
	comments, total, err := mysql_logic.ListCommentByPostId(ctx.Request.Context(), postId, pageNo, pageSize)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	// 提取所有用户ID
	userIds := make([]int64, 0, len(comments))
	userMap := make(map[int64]bool)
	for _, comment := range comments {
		if !userMap[comment.UserId] {
			userMap[comment.UserId] = true
			userIds = append(userIds, comment.UserId)
		}
	}

	// 批量查询作者信息
	authorMap := make(map[int64]*mysql_model.Users)
	if len(userIds) > 0 {
		authorMap, err = mysql_logic.GetUsersByIds(ctx.Request.Context(), userIds)
		if err != nil {
			resp.Code = utils.FailCode
			resp.Message = err.Error()
			return resp
		}
	}

	// 构建响应
	resp.Total = total
	resp.CommentList = make([]*pb.Comment, len(comments))
	for i, comment := range comments {
		author := authorMap[comment.UserId]
		authorPb := &pb.User{
			UserId: strconv.FormatInt(comment.UserId, 10),
		}
		if author != nil {
			authorPb.Username = author.Username
			authorPb.Email = author.Email
		}

		resp.CommentList[i] = &pb.Comment{
			CommentId:  strconv.FormatInt(comment.Id, 10),
			Content:    comment.Content,
			PostId:     strconv.FormatInt(comment.PostId, 10),
			CreateTime: comment.CreatedAt,
			UpdateTime: comment.UpdatedAt,
			Author:     authorPb,
		}
	}

	return resp
}
