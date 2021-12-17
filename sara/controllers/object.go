package controllers

import (
	"encoding/json"
	"sara/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectemail := models.AddOne(ob)
	o.Data["json"] = map[string]string{"Email": objectemail}
	o.ServeJSON()
}
