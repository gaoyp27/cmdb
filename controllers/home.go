package controllers

import (
	"cmdb/base/controllers/auth"
)

// HomeController 首页控制器
type HomeController struct {
	auth.AuthorizationController
}

// Index  首页显示方法
func (c *HomeController) Index() {
	// session 检查
	c.TplName = "home/index.html"
}
