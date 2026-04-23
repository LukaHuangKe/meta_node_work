package handler

import (
	"phase1/phase1_work/mysql/mysql_model"
	"strconv"
	"time"

	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req pb.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	//c.JSON(200, handleRegister(c.Request.Context(), &req))
	c.JSON(200, handleRegister(c, &req))
}

func handleRegister(ctx *gin.Context, req *pb.RegisterReq) (resp *pb.RegisterResp) {
	resp = &pb.RegisterResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	// 参数校验
	if len(req.GetUsername()) == 0 || len(req.GetPassword()) == 0 || len(req.GetEmail()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = utils.FailCodeMsg
		return resp
	}

	// 写入DB
	now := time.Now().UnixMilli()
	user := &mysql_model.Users{
		Username:  req.GetUsername(),
		Password:  req.GetPassword(),
		Email:     req.GetEmail(),
		Status:    1,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: 0,
	}

	userId, err := mysql_logic.CreateUser(ctx, user)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	user.Id = userId
	resp.User = &pb.User{
		UserId:   strconv.FormatInt(userId, 10),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
	}

	// 返回token
	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}
	resp.JwtToken = token
	return resp
}
