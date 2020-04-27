package users

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	usersTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := usersTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("个人简介", "personal", db.Text)
	info.AddField("联系我", "contact", db.Text)
	info.AddField("关于", "about", db.Text)

	info.SetTable("users").SetTitle("个人简介").SetDescription("")

	formList := usersTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("个人简介", "personal", db.Text, form.RichText)
	formList.AddField("联系我", "contact", db.Text, form.RichText)
	formList.AddField("关于", "about", db.Text, form.RichText)

	formList.SetTable("users").SetTitle("个人简介").SetDescription("")

	return usersTable
}
