/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 18:45
 */

package controller

const (
	//Success
	OK = 0

	//Error
	ErrorInvalidMethodOfRequest = iota + 10000
	ErrorInvalidParamOfInterface
	ErrorInvalidUser
	ErrorInvalidEmailOrMobile
	ErrorDatabase

	ErrorBadFormatOfEmailOrMobile
	ErrorBadFormatOfAlias

	ErrorOfServer
)


//CustomStatusText 自定义状态码说明
var CustomStatusText = map[int]string{
	//Success
	OK: "Success",

	//Error
	ErrorInvalidMethodOfRequest:  "无效的请求方法",
	ErrorInvalidParamOfInterface: "无效的接口参数",
	ErrorInvalidUser:             "无效的用户",
	ErrorInvalidEmailOrMobile:    "无效的邮箱或手机",
	ErrorDatabase:               "数据库操作错误",

	ErrorBadFormatOfEmailOrMobile: "邮箱或手机号格式错误",
	ErrorBadFormatOfAlias:         "用户名/昵称格式错误",
	ErrorOfServer: "服务器错误",

}

//FrontData 返回前端的数据
type FrontData struct {
	Code    int
	Message string
	Data    interface{}
}

//NewFrontData 初始化返回前端的数据
func NewFrontData(code int, data interface{}) *FrontData {
	return &FrontData{
		Code:    code,
		Message: CustomStatusText[code],
		Data:    data,
	}
}

//NewServerError ...
func NewServerError(err error) *FrontData {
	return &FrontData{
		Code:    ErrorOfServer,
		Message: err.Error(),
		Data:    nil,
	}
}