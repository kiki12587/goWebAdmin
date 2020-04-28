package article_message

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleMessageTable(ctx *context.Context) table.Table {

	articleMessageTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleMessageTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("文章ID", "article_id", db.Int)
	info.AddField("留言昵称", "nickname", db.Char)
	info.AddField("邮箱", "email", db.Varchar)
	info.AddField("留言内容", "message", db.Varchar)
	info.AddField("留言时间", "create_at", db.Datetime)

	info.SetTable("article_message").SetTitle("用户留言").SetDescription("")

	formList := articleMessageTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Article_id", "article_id", db.Int, form.Number)
	formList.AddField("Nickname", "nickname", db.Char, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Message", "message", db.Varchar, form.Text)
	formList.AddField("Create_at", "create_at", db.Datetime, form.Datetime)

	formList.SetTable("article_message").SetTitle("用户留言").SetDescription("")

	return articleMessageTable
}
