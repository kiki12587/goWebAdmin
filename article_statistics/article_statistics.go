package article_statistics

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleStatisticsTable(ctx *context.Context) table.Table {

	articleStatisticsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleStatisticsTable.GetInfo()
	info.HideNewButton()
	info.HideExportButton()
	info.HideDetailButton()
	info.SetSortDesc()
	info.AddField("文章ID", "id", db.Int).FieldFilterable()
	info.AddField("文章评论数量", "article_comment_num", db.Int)
	info.AddField("文章浏览量", "article_browse_num", db.Int)

	info.SetTable("article_statistics").SetTitle("文章数据统计").SetDescription("")

	formList := articleStatisticsTable.GetForm()

	formList.AddField("文章ID", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("文章评论数量", "article_comment_num", db.Int, form.Number)
	formList.AddField("文章浏览量", "article_browse_num", db.Int, form.Number)

	formList.SetTable("article_statistics").SetTitle("文章数据统计").SetDescription("")

	return articleStatisticsTable
}
