package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"skrshop-api/core"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		e := c.Errors.Last()
		err := e.Err
		var errStr string
		//// TODO::记录错误日志
		if e != nil {
			switch err.(type) {
			case validator.ValidationErrors:
				errStr = core.Translate(err.(validator.ValidationErrors))
			case *json.UnmarshalTypeError:
				unmarshalTypeError := err.(*json.UnmarshalTypeError)
				errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
			default:
				errStr = e.Err.Error()
			}
		}
		core.ErrorParamsResp(c, errStr)
	}

}
