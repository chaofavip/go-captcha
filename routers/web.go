package routers

import (
	captcha "captchaCode/http/controller/captcha"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	// 创建一个验证码路由
	verifyCode := router.Group("captcha")
	{
		// 验证码业务，该业务无需专门校验参数，所以可以直接调用控制器
		verifyCode.GET("/", (&captcha.Captcha{}).GenerateId)                 //  获取验证码ID
		verifyCode.GET("/:captchaId", (&captcha.Captcha{}).GetImg)           // 获取图像地址
		verifyCode.GET("/:captchaId/:value", (&captcha.Captcha{}).CheckCode) // 校验验证码
	}
	return router
}
