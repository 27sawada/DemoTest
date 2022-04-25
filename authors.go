package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthorsTable(ctx *context.Context) table.Table {

	authors := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := authors.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("First_name", "first_name", db.Varchar)
	info.AddField("Last_name", "last_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Birthdate", "birthdate", db.Date)
	info.AddField("Added", "added", db.Timestamp)

	info.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	formList := authors.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("First_name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last_name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Birthdate", "birthdate", db.Date, form.Datetime)
	formList.AddField("Added", "added", db.Timestamp, form.Datetime)

	formList.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	return authors
}
