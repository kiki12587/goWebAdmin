package article

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleTable(ctx *context.Context) table.Table {

	articleTable := table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true, // 是否可以新增
		Editable:   true, // 是否可以编辑
		Deletable:  true, // 是否可以删除
		Exportable: true, // 是否可以导出为excel
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	},
	)

	info := articleTable.GetInfo()
	info.AddField("文章ID", "id", db.Int).FieldFilterable()
	info.AddField("文章标题", "title", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

	//info.HideExportButton()
	//info.HideFilterArea()
	//info.SetSortDesc()
	info.SetSortField("id")

	//info.AddField("文章分类", "category_id", db.Int)

	info.AddField("文章分类", "category_name", db.Varchar).FieldJoin(types.Join{
		Table:     "article_category", // 连表的表名
		Field:     "category_id",      // 要连表的字段
		JoinField: "id",               // 连表的表的字段
	})

	info.AddField("文章标题", "title", db.Varchar)
	info.AddField("文章作者", "auth", db.Char)
	info.AddField("文章简介", "info", db.Varchar)
	//info.AddField("文章头图", "image", db.Varchar)

	info.AddField("文章头图", "image", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value != "" {
			return "<img alt='Responsive image' class='img-rounded' src=/uploads/" + model.Value + ">"
		}
		return "<span class='label btn btn-info'>" + "暂未上传" + "</span>"
	})

	info.AddField("文章标签", "label", db.Varchar)
	//info.AddField("是否为精选文章", "selected_articles_status", db.Enum)
	info.AddField("是否为精选文章", "selected_articles_status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "<span class='label btn btn-danger'>" + "否" + "</span>"
		}
		if model.Value == "1" {
			return "<span class='label btn btn-success'>" + "是" + "</span>"
		}
		return "<span class='label btn btn-danger'>" + "否" + "</span>"
	})

	//info.AddField("是否上架", "article_status", db.Enum)
	info.AddField("是否上架", "article_status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "<span class='label btn btn-danger'>" + "否" + "</span>"
		}
		if model.Value == "1" {
			return "<span class='label btn btn-success'>" + "是" + "</span>"
		}
		return "<span class='label btn btn-danger'>" + "否" + "</span>"
	})

	info.SetTable("article").SetTitle("文章列表").SetDescription("")

	formList := articleTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()

	formList.AddField("文章分类", "category_id", db.Int, form.Number)

	formList.AddField("文章标题", "title", db.Varchar, form.Text)
	formList.AddField("文章作者", "auth", db.Char, form.Text)
	formList.AddField("文章简介", "info", db.Varchar, form.Text)
	formList.AddField("文章头图", "image", db.Varchar, form.File)
	formList.AddField("文章内容", "content", db.Text, form.RichText)
	formList.AddField("文章标签", "label", db.Varchar, form.Text)

	// 使用 FieldOptions 设置 radio 类型内容
	formList.AddField("是否为精选文章", "selected_articles_status", db.Enum, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "0"},
			{Text: "是", Value: "1"},
		}).FieldDefault("0")

	//formList.AddField("是否为精选文章", "selected_articles_status", db.Enum, form.Text)

	formList.AddField("是否上架", "article_status", db.Enum, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "0"},
			{Text: "是", Value: "1"},
		}).FieldDefault("1")

	formList.AddField("发布时间", "create_at", db.Timestamp, form.Datetime)
	formList.AddField("更新时间", "update_at", db.Timestamp, form.Datetime)

	//formList.AddField("是否上架", "article_status", db.Enum, form.Text)

	formList.SetTable("article").SetTitle("文章属性").SetDescription("")

	return articleTable
}
