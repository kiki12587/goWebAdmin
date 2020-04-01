package article_category

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleCategoryTable(ctx *context.Context) table.Table {

	articleCategoryTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleCategoryTable.GetInfo()
	//info.HideExportButton()
	//info.HideFilterArea()
	//info.SetSortDesc()

	info.SetSortField("create_at")

	info.AddField("文章分类ID", "id", db.Int).FieldFilterable()
	info.AddField("文章分类名称", "category_name", db.Char).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

	//info.AddField("Category_name", "category_name", db.Char)
	//info.AddField("Create_at", "create_at", db.Datetime)
	//info.AddField("Status", "status", db.Enum)
	//
	//info.AddField("文章分类名称", "category_name", db.Char)

	//info.AddField("文章分类是否显示", "status", db.Enum)
	info.AddField("文章分类是否显示", "status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "<span class='label btn btn-danger'>" + "否" + "</span>"
		}
		if model.Value == "1" {
			return "<span class='label btn btn-success'>" + "是" + "</span>"
		}
		return "<span class='label btn btn-danger'>" + "否" + "</span>"
	})
	info.AddField("创建时间", "create_at", db.Datetime)
	info.SetTable("article_category").SetTitle("文章分类").SetDescription("")

	formList := articleCategoryTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("文章分类名称", "category_name", db.Char, form.Text)
	formList.AddField("发布时间", "create_at", db.Timestamp, form.Datetime)
	//formList.AddField("Status", "status", db.Enum, form.Text)
	formList.AddField("分类是否上架", "status", db.Enum, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "0"},
			{Text: "是", Value: "1"},
		}).FieldDefault("1")

	formList.SetTable("article_category").SetTitle("文章分类").SetDescription("")

	return articleCategoryTable
}
