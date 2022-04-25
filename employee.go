package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEmployeeTable(ctx *context.Context) table.Table {

	employee := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := employee.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint)
	info.AddField("Department", "department", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Job", "job", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("employee").SetTitle("Employee").SetDescription("Employee")

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
