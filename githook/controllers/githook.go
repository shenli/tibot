package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"github.com/ngaut/log"
	"github.com/shenli/tibot/githook/hook"
	"github.com/shenli/tibot/githook/models"
	"github.com/shenli/tibot/slack"
)

// Operations about object
type GithookController struct {
	beego.Controller
	Git *hook.Hook
}

func (o *GithookController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	o.Controller.Init(ct, controllerName, actionName, app)
	o.Git = &hook.Hook{}

	// Config
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("Parse ini fail: ", err)
	}

	cfg := iniconf.String("SlackConfig")
	fmt.Println("Get slack config: ", cfg)
	s, err := slack.NewSlack(cfg)
	if err != nil {
		fmt.Println("Init slack client meet error: ", err)
	}
	o.Git.SetSlackClient(s)
}

// @Title Create
// @Description create object
// @Param	body		body 	models.GitGithook	true		"The object content"
// @Success 200 {string} models.GitGithook.Id
// @Failure 403 body is empty
// @router / [post]
func (o *GithookController) Post() {
	// json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	event := o.Ctx.Input.Header("x-github-event")
	err := o.Git.HandleEvent(event, o.Ctx.Input.RequestBody)
	if err != nil {
		log.Errorf("Parse event %s meet error %s", event, err)
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.GitGithook
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *GithookController) Get() {
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
