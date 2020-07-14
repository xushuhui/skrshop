package api

import (
	"github.com/gin-gonic/gin"
)

type UserToken struct {
	Uid string
}

func Login(c *gin.Context) (resp interface{}, err error) {
	//var req request.SignIn
	//if err = core.ParseRequest(c, &req); err != nil {
	//	return
	//}
	//
	//data, err := model.GetStaffInfoOne("phone=?", req.Phone)
	//// 正确密码验证
	//err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	//if err != nil {
	//	core.FailResp(c, code.ErrorPassWord)
	//}
	return &UserToken{Uid: string(1)}, nil

}
