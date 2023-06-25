package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/models"
	"fmt"
)

type UserController struct {
	auth.AuthorizationController
}

func (c *UserController) Query() {
	fmt.Println("Query")
	q := c.GetString("q")

	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/query.html"
}
