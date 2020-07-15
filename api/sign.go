package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"skrshop-api/code"
	"skrshop-api/core"
	"skrshop-api/model"
	"skrshop-api/request"
	"skrshop-api/utils"
)

type UserToken struct {
	Uid string
}

func Login(c *gin.Context) {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		return
	}

	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(c, code.ErrorPassWord)
	}
	data, err := utils.GenerateToken(userModel.Id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
func PhoneLogin(c *gin.Context) {

}
func MiniappLogin() {

}
