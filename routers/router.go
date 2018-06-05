package routers

import (
	"github.com/astaxie/beego"
	"admin/controllers/common"
	//"admin/controllers"
	"admin/controllers/rbac"
)

func init() {
	
	//beego.Router("/",&controllers.MainController{})
	beego.Router("/",&common.MainController{},"*:Admin")
	beego.Router("public/index",&common.MainController{},"*:Index")
	beego.Router("public/login",&common.MainController{},"*:Login")
	beego.Router("public/logout",&common.MainController{},"*:Logout")
	beego.Router("public/changepwd",&common.MainController{},"*:Changepwd")
	
	
	
	beego.Router("admin/user?:page",&rbac.UserController{},"*:Index")
	beego.Router("admin/user/add",&rbac.UserController{},"*:UserAdd")
	beego.Router("admin/user/adds",&rbac.UserController{},"*:UserAdds")
	beego.Router("admin/user/edit",&rbac.UserController{},"*:UserEdit")
	beego.Router("admin/user/edits",&rbac.UserController{},"*:UserEdits")
	
	
	beego.Router("admin/role/?:page",&rbac.RoleController{},"*:Index")
	beego.Router("admin/role/edit/:id",&rbac.RoleController{},"*:RoleEdit")
	beego.Router("admin/role/edits",&rbac.RoleController{},"*:RoleEdits")
	beego.Router("admin/role/add",&rbac.RoleController{},"*:RoleAdd")
	beego.Router("admin/role/adds",&rbac.RoleController{},"*:RoleAdds")
	
	
	
	
	//
	//beego.Router("/", &rbac.MainController{}, "*:Index")
	//beego.Router("/public/index", &rbac.MainController{}, "*:Index")
	//beego.Router("/public/login", &rbac.MainController{}, "*:Login")
	//beego.Router("/public/logout", &rbac.MainController{}, "*:Logout")
	//beego.Router("/public/changepwd", &rbac.MainController{}, "*:Changepwd")
	//
	//beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	//beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	//beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	//beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")
	//
	//beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	//beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	//beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")
	//
	//beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	//beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	//beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	//beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")
	//
	//beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	//beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	//beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	//beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	//beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	//beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	//beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	//beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")
}

