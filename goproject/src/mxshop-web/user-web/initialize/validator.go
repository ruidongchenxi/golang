package initialize

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)
var trans ut.Translator
func InintTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义来
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文
		enT := en.New() //英文翻译
		//第一个参数备用语言环境，后续参数应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		if trans, ok = uni.GetTranslator(locale); !ok {
			return fmt.Errorf("GetTranslator(%s)", locale)
		} else {
			switch locale {
			case "en":
				enTranslations.RegisterDefaultTranslations(v, trans)
			case "zh":
				zhTranslations.RegisterDefaultTranslations(v, trans)
			default:
				enTranslations.RegisterDefaultTranslations(v, trans)

			}
			return
		}
		//return

	}
	return
}