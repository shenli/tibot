package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/shenli/tibot/githook/models"
	"github.com/shenli/tibot/slack"
)

// Operations about object
type GithookController struct {
	beego.Controller
	slackClient *slack.Slack
}

// @Title Create
// @Description create object
// @Param	body		body 	models.GitGithook	true		"The object content"
// @Success 200 {string} models.GitGithook.Id
// @Failure 403 body is empty
// @router / [post]
func (o *GithookController) Post() {
	// json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	fmt.Printf("Get Post: %s\n", o.Ctx.Input.RequestBody)
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.GitGithook
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *GithookController) Get() {
	fmt.Println("Get One")
	ob := &models.Githook{
		GithookId:  "xxx",
		Score:      100,
		PlayerName: "shenli",
	}
	o.Data["json"] = ob
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.GitGithook
// @Failure 403 :objectId is empty
// @router / [get]
func (o *GithookController) GetAll() {
	fmt.Println("Get All")
	obs := models.GitGetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.GitGithook	true		"The body"
// @Success 200 {object} models.GitGithook
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *GithookController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Githook
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.GitUpdate(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *GithookController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.GitDelete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
