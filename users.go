package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	//editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := users.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable().FieldSortable().SetSortAsc()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "男"
		}
		if model.Value == "1" {
			return "女"
		}
		return "保密"

		// }).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		// 	{Value: "0", Text: "男"},
		// 	{Value: "1", Text: "女"},
		// }).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		// 	{Value: "0", Text: "men"},
		// 	{Value: "1", Text: "women"},
	})
	info.AddField("City", "city", db.Varchar)
	info.AddField("Ip", "ip", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Number)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Ip", "ip", db.Varchar, form.Ip)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
