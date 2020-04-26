package article_user_message

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleUserMessageTable(ctx *context.Context) table.Table {

	articleUserMessageTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleUserMessageTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("留言内容", "message", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("留言人的昵称", "nickname", db.Char).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

	info.AddField("留言人的昵称", "nickname", db.Char)
	info.AddField("内容", "message", db.Varchar)
	info.AddField("邮箱", "email", db.Varchar)
	info.AddField("时间", "create_at", db.Datetime)

	info.SetTable("article_user_message").SetTitle("用户留言").SetDescription("")

	formList := articleUserMessageTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()

	formList.AddField("留言人的昵称", "nickname", db.Char, form.Text)
	formList.AddField("内容", "message", db.Varchar, form.Text)
	formList.AddField("邮箱", "email", db.Varchar, form.Email)
	formList.AddField("时间", "create_at", db.Datetime, form.Datetime)

	formList.SetTable("article_user_message").SetTitle("用户留言").SetDescription("")

	return articleUserMessageTable
}
