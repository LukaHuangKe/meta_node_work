package handler

import (
	"strconv"

	"phase1/phase1_work/mysql/mysql_logic"
	"phase1/phase1_work/pb"
	"phase1/phase1_work/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": "400", "message": "参数错误"})
		return
	}

	c.JSON(200, handleLogin(c, &req))
}

func handleLogin(c *gin.Context, req *pb.LoginReq) (resp *pb.LoginResp) {
	resp = &pb.LoginResp{
		Code:    utils.SuccessCode,
		Message: utils.SuccessCodeMsg,
	}

	if len(req.GetUsername()) == 0 || len(req.GetPassword()) == 0 {
		resp.Code = utils.FailCode
		resp.Message = utils.FailCodeMsg
		return resp
	}

	user, err := mysql_logic.GetUserByUsernameAndPassword(c.Request.Context(), req.GetUsername(), req.GetPassword())
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}

	resp.User = &pb.User{
		UserId:   strconv.FormatInt(user.Id, 10),
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		resp.Code = utils.FailCode
		resp.Message = err.Error()
		return resp
	}
	resp.JwtToken = token
	return resp
}
