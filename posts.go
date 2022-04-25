package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPostsTable(ctx *context.Context) table.Table {

	posts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := posts.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Author_id", "author_id", db.Int)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Content", "content", db.Text)
	info.AddField("Date", "date", db.Date)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := posts.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Author_id", "author_id", db.Int, form.Number)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Text, form.RichText)
	formList.AddField("Date", "date", db.Date, form.Datetime)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return posts
}
