package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"skrshop-api/app/code"
	"skrshop-api/app/model"
	"skrshop-api/app/request"
	"skrshop-api/core"
	"skrshop-api/lib"
)

type UserToken struct {
	Uid string
}

func Login(c *gin.Context) {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	if err != nil {
		c.Error(err)
		return
	}
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(c, code.ErrorPassWord)
	}
	data, err := lib.GenerateToken(userModel.Id)
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
