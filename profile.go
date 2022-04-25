package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProfileTable(ctx *context.Context) table.Table {

	profile := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := profile.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Uuid", "uuid", db.Varchar)
	info.AddField("Photos", "photos", db.Varchar)
	info.AddField("Resume", "resume", db.Varchar)
	info.AddField("Resume_size", "resume_size", db.Int)
	info.AddField("Finish_state", "finish_state", db.Tinyint)
	info.AddField("Finish_progress", "finish_progress", db.Int)
	info.AddField("Pass", "pass", db.Tinyint)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("profile").SetTitle("Profile").SetDescription("Profile")

	formList := profile.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Uuid", "uuid", db.Varchar, form.Text)
	formList.AddField("Photos", "photos", db.Varchar, form.Text)
	formList.AddField("Resume", "resume", db.Varchar, form.Text)
	formList.AddField("Resume_size", "resume_size", db.Int, form.Number)
	formList.AddField("Finish_state", "finish_state", db.Tinyint, form.Number)
	formList.AddField("Finish_progress", "finish_progress", db.Int, form.Number)
	formList.AddField("Pass", "pass", db.Tinyint, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("profile").SetTitle("Profile").SetDescription("Profile")

	return profile
}
