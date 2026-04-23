package handler

import (
	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/mysql/mysql_model"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPost(c *gin.Context) {
	var req pb.ListPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleListPost(c, &req))
}

func handleListPost(ctx *gin.Context, req *pb.ListPostReq) (resp *pb.ListPostResp) {
	resp = &pb.ListPostResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
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

	// 查询帖子列表
	posts, total, err := mysql_logic.ListPost(
		ctx.Request.Context(),
		pageNo,
		pageSize,
		req.GetStartTime(),
		req.GetEndTime(),
		req.GetPostId(),
	)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	// 提取所有用户ID
	userIds := make([]int64, 0, len(posts))
	userMap := make(map[int64]bool)
	for _, post := range posts {
		if !userMap[post.UserId] {
			userMap[post.UserId] = true
			userIds = append(userIds, post.UserId)
		}
	}

	// 批量查询作者信息
	authorMap := make(map[int64]*mysql_model.Users)
	if len(userIds) > 0 {
		var err error
		authorMap, err = mysql_logic.GetUsersByIds(ctx.Request.Context(), userIds)
		if err != nil {
			resp.Code = utils.FailCode
			resp.Message = err.Error()
			return resp
		}
	}

	// 构建响应
	resp.Total = total
	resp.PostList = make([]*pb.Post, len(posts))
	for i, post := range posts {
		author := authorMap[post.UserId]
		authorPb := &pb.User{
			UserId: strconv.FormatInt(post.UserId, 10),
		}
		if author != nil {
			authorPb.Username = author.Username
			authorPb.Email = author.Email
		}

		resp.PostList[i] = &pb.Post{
			PostId:     strconv.FormatInt(post.Id, 10),
			Title:      post.Title,
			Content:    post.Content,
			CreateTime: post.CreatedAt,
			UpdateTime: post.UpdatedAt,
			Author:     authorPb,
		}
	}

	return resp
}
