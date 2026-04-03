package model
//根据业务逻辑需要，自定义一些错误
import(
	"errors"
)
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS = errors.New("用户ID存在")
	ERROR_USER_PWD = errors.New("密码错误")


)