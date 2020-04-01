package article_comment

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleCommentTable(ctx *context.Context) table.Table {

	articleCommentTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleCommentTable.GetInfo()
	info.HideNewButton()
	info.HideExportButton()
	info.HideDetailButton()
	info.SetSortDesc()

	info.AddField("文章ID", "id", db.Int).FieldFilterable()
	info.AddField("评论内容", "article_comment_content", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("回复内容", "article_comment_reply_content", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

	info.AddField("文章评论人名称", "article_comment_name", db.Char)
	info.AddField("评论内容", "article_comment_content", db.Varchar)
	info.AddField("回复内容", "article_comment_reply_content", db.Varchar)

	//info.AddField("评论是否显示", "article_comment_status", db.Enum)

	info.AddField("评论是否显示", "article_comment_status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "<span class='label btn btn-danger'>" + "否" + "</span>"
		}
		if model.Value == "1" {
			return "<span class='label btn btn-success'>" + "是" + "</span>"
		}
		return "<span class='label btn btn-danger'>" + "否" + "</span>"
	})

	//info.AddField("回复是否显示", "article_comment_reply_status", db.Enum)

	info.AddField("回复是否显示", "article_comment_reply_status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "<span class='label btn btn-danger'>" + "否" + "</span>"
		}
		if model.Value == "1" {
			return "<span class='label btn btn-success'>" + "是" + "</span>"
		}
		return "<span class='label btn btn-danger'>" + "否" + "</span>"
	})

	info.AddField("评论人邮箱", "article_comment_email", db.Char)
	info.AddField("评论时间", "create_at", db.Timestamp)

	info.SetTable("article_comment").SetTitle("评论属性").SetDescription("")

	formList := articleCommentTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	//formList.AddField("Article_comment_name", "article_comment_name", db.Char, form.Text)
	//formList.AddField("Article_comment_content", "article_comment_content", db.Varchar, form.Text)
	//formList.AddField("Article_comment_reply_status", "article_comment_reply_status", db.Enum, form.Text)

	formList.AddField("回复是否显示", "Article_comment_reply_status", db.Enum, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "0"},
			{Text: "是", Value: "1"},
		}).FieldDefault("0")

	//formList.AddField("Article_comment_reply_content", "article_comment_reply_content", db.Varchar, form.Text)

	formList.AddField("回复是否显示", "Article_comment_reply_status", db.Enum, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "0"},
			{Text: "是", Value: "1"},
		}).FieldDefault("0")

	//formList.AddField("回复是否显示", "Article_comment_reply_status", db.Enum, form.Text)
	//formList.AddField("Create_at", "create_at", db.Timestamp, form.Datetime)
	//formList.AddField("Article_comment_email", "article_comment_email", db.Char, form.Text)

	formList.SetTable("article_comment").SetTitle("评论属性").SetDescription("Article_comment")

	return articleCommentTable
}
