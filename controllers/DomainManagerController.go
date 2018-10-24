package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lhtzbj12/sdrms/enums"
	"github.com/lhtzbj12/sdrms/models"
	"strconv"
	"strings"
)

// DomainManagerController 接入域名管理控制器
type DomainManagerController struct {
	BaseController
}

// Index 接口首页
func (c *DomainManagerController) Index() {
	// 是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	// 将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "domainmanager/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "domainmanager/index_footerjs.html"
	// 页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("DomainManagerController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("DomainManagerController", "Delete")
}

// Edit 添加、编辑角色界面
func (c *DomainManagerController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := models.DomainManager{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m

	c.setTpl("domainmanager/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "domainmanager/edit_footerjs.html"
}

// Save 添加、编辑页面 保存
func (c *DomainManagerController) Save() {
	var err error
	m := models.DomainManager{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}

	} else {
		if _, err = o.Update(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}

}

// DataGrid 数据接口管理首页 表格获取数据
func (c *DomainManagerController) DataGrid() {
	var params models.DomainManagerQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	data, total := models.DomainManagerPageList(&params)
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//Delete 批量删除
func (c *DomainManagerController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.DomainManagerBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
