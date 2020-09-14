package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

var MessageTmpls = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小值 为 %d",
	"Max":          "最大值 为 %d",
	"Range":        "范围 为 %d 到 %d",
	"MinSize":      "最短长度 为 %d",
	"MaxSize":      "最大长度 为 %d",
	"Length":       "长度必须 为 %d",
	"Alpha":        "必须是有效的字母",
	"Numeric":      "必须是有效的数字",
	"AlphaNumeric": "必须是有效的字母或数字",
	"Match":        "必须匹配 %s",
	"NoMatch":      "必须不匹配 %s",
	"AlphaDash":    "必须是有效的字母、数字或连接符号(-_)",
	"Email":        "必须是有效的电子邮件地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "必须是有效的base64字符",
	"Mobile":       "必须是有效的手机号码",
	"Tel":          "必须是有效的电话号码",
	"Phone":        "必须是有效的电话或移动电话号码",
	"ZipCode":      "必须是有效的邮政编码",
}

func BindAndValid(c *gin.Context, inputData interface{}) error {
	err := c.ShouldBind(inputData)
	if err != nil {
		return err
	}

	validation.SetDefaultMessage(MessageTmpls)
	valid := validation.Validation{}
	check, err := valid.Valid(inputData)
	if err != nil {
		return err
	}

	if !check {
		return valid.Errors[0]
	}

	return nil
}
