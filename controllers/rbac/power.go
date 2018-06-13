/**
 * User: yp
 * Time: 2018/6/11  下午4:58
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type PowerController struct {
	common.CommonController
}

func (this *PowerController) Index() {
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	powers1,powers2,powers3 := models.GroupList()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	this.TplName = this.GetTemplatetype() + "/powerlist.html"
}

// 管理
func (this *PowerController) PowerEdit(){
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	powerId,_ := this.GetInt64(":id")
	powers := models.GetPowerId(powerId)
	powers1,powers2,powers3 := models.GroupList()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	this.Data["powers"] = &powers
	this.TplName = this.GetTemplatetype() + "/poweredit.html"
}
func (this *PowerController) PowerEdits(){
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	powername := this.GetString("powername")
	controller := this.GetString("controller")
	action := this.GetString("action")
	pid,_ := this.GetInt("pid")
	//var lpid int64
	//lpid = int64(pid)
	id,_ := this.GetInt64("id")
	powerid,_ := this.GetInt64("powerid")
	status,_ := this.GetInt("status")
	o := orm.NewOrm()
	power1 := models.Power{}
	power2 := models.Power{}
	power := models.Power{}
	power1.Id = id
	power2.Id = powerid
	o.Read(&power1) //选择的
	o.Read(&power2) //当前的
	o.Read(&power) //当前的
	if pid == 0 { //顶级
		power.Level = 2
		power.Pid = int(power1.Id)
	}else{
		if id == powerid {
			power.Level = power1.Level
			power.Pid = power1.Pid
		}else {
			//如果选择的等级大于现在的等级
			if power1.Level == 1 &&  power2.Level == 1 {
				power.Level = 2
			}else if power1.Level == 2 &&  power2.Level == 2 {
				power.Level = 3
			}else if power1.Level == 1 && power2.Level == 2 {
				power.Level = 2
			}else if power1.Level == 1 && power2.Level == 3 {
				power.Level = 2
			}else if power1.Level == 2 && power2.Level == 3 {
				power.Level = 3
			}else if power1.Level == 2 && power2.Level == 1 {
				power.Level = 3
			}
			power.Pid = int(power1.Id)
		}
	}
	power.Controller = controller
	power.Action = action
	power.Status = status
	power.Powername = powername
	power.Id = powerid
	power.Updatetime = time.Now()
	power.Createtime = time.Now()
	beego.Info(power)
	if num, err := o.Update(&power); err == nil {
		this.Rsp(true,"修改成功")
		return
	}else{
		beego.Info(num)
		this.Rsp(false,"修改失败")
		return
	}
	
}

func (this *PowerController) PowerAdd() {
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	tree := this.GetTree()
	powers1,powers2,powers3 := models.GroupList()
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	this.Data["tree"] = &tree
	//this.Data["users"] = userinfo
	this.TplName = this.GetTemplatetype() + "/poweradd.html"
}

func (this *PowerController) PowerAdds() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	id,_ := this.GetInt64("id")
	powername := this.GetString("powername")
	controller := this.GetString("controller")
	action := this.GetString("action")
	pid,_ := this.GetInt("pid")
	status,_ := this.GetInt("status")
	o := orm.NewOrm()
	power := models.Power{}
	if pid == 0 { //顶级
		power.Level = 1
		power.Pid = pid
	}else{
		power.Id = id
		o.Read(&power)
		if power.Level == 1 {
			power.Level = 2
		}else {
			power.Level = 3
		}
		power.Pid = int(power.Id)
	}
	power.Controller = controller
	power.Action = action
	power.Status = status
	power.Powername = powername
	if _, err := o.Insert(&power); err == nil {
		this.Rsp(true,"添加成功")
		return
	}
	this.Rsp(false,"添加失败")
	return
}