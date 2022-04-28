package main

import (
	"github.com/GoAdminGroup/components/login"
	_ "github.com/GoAdminGroup/components/login/theme2"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // 适配器
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql 驱动
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/sword" // ui主题

	"io/ioutil"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	template.AddComp(chartjs.NewChart())

	login.Init(login.Config{
		Theme:         "theme2",
		CaptchaDigits: 5,
	})

	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(Generators).
		Use(r); err != nil {
		panic(err)
	}

	// 载入对应验证码驱动，如没使用不用载入
	adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})
	// captcha.Add(login.CaptchaDriverKeyDefault, new(login.DigitsCaptcha))

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
