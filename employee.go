package main

import (
	template2 "html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEmployeeTable(ctx *context.Context) table.Table {

	employee := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := employee.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable().FieldSortable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "0" {
			return "男"
		}
		return "女"
	})
	info.AddField("Department", "department", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Job", "job", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp).FieldSortable()

	department := ctx.Query("department")

	info.SetTable("employee").SetTitle("Employee").SetDescription("Employee").
		SetWrapper(func(content template2.HTML) template2.HTML {
			col1 := `<div style="margin-left:243px;">` + content + `</div>`

			tree := template.Default().TreeView().SetTree(types.TreeViewData{
				Data: types.TreeViewItems{
					{
						Text: "TLZKJ",
						Href: "/admin/info/employee?__go_admin_no_animation_=true",
						Nodes: types.TreeViewItems{
							{
								Text: "技术",
								State: types.TreeViewItemState{
									Expanded: department == "前端" || department == "中台" || department == "后端",
								},
								Nodes: types.TreeViewItems{
									{
										Text: "前端",
										Href: "/admin/info/employee?department=前端&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "前端",
											Selected: department == "前端",
										},
									}, {
										Text: "中台",
										Href: "/admin/info/employee?department=中台&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "中台",
											Selected: department == "中台",
										},
									}, {
										Text: "后端",
										Href: "/admin/info/employee?department=后端&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "后端",
											Selected: department == "后端",
										},
									},
								},
							},
							{
								Text: "管理",
								State: types.TreeViewItemState{
									Expanded: department == "销售" || department == "前台",
								},
								Nodes: types.TreeViewItems{
									{
										Text: "销售",
										Href: "/admin/info/employee?department=销售&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "销售",
											Selected: department == "销售",
										},
									}, {
										Text: "前台",
										Href: "/admin/info/employee?department=前台&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "前台",
											Selected: department == "前台",
										},
									},
								},
							}, {
								Text: "人力",
								Href: "/admin/info/employee?department=人力&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "人力",
									Selected: department == "人力",
								},
							},
						},
					},
				},
				ExpandIcon:        "fa fa-angle-right",
				CollapseIcon:      "fa fa-angle-down",
				SelectedBackColor: "#fbfbfb",
				SelectedColor:     "#333333",
				EnableLinks:       true,
			}).GetContent()

			col2 := `<div style="position: absolute;width:230px;">` + template.Default().Box().SetHeader("组织结构").
				WithHeadBorder().SetBody(tree).GetContent() + `</div>`
			return `<div style="width:100%;">` + col2 + col1 + `</div>`
		})

	formList := employee.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Number)
	formList.AddField("Department", "department", db.Varchar, form.Text)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Job", "job", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("employee").SetTitle("Employee").SetDescription("Employee")

	return employee
}
