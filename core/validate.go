package core

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

func isMobile(fl validator.FieldLevel) bool {
	mobile := strconv.Itoa(int(fl.Field().Uint()))
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}

var trans ut.Translator

func InitValidate() {
	// 中文翻译
	zh := zhongwen.New()

	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json tag的自定义方法
		//v.RegisterTagNameFunc(jsonTag)
		// 验证器注册翻译器
		zhtranslations.RegisterDefaultTranslations(v, trans)
		// 自定义验证方法
		v.RegisterValidation("isMobile", isMobile)
	}
}
func jsonTag(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name

}
func Translate(errs validator.ValidationErrors) string {
	var errList []string

	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}
