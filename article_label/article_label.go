package article_label

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleLabelTable(ctx *context.Context) table.Table {

	articleLabelTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := articleLabelTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("标签名称", "label_name", db.Varchar)

	info.SetTable("article_label").SetTitle("标签库").SetDescription("")

	formList := articleLabelTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()

	formList.AddField("标签名称", "label_name", db.Varchar, form.Text)

	formList.SetTable("article_label").SetTitle("标签库").SetDescription("")

	return articleLabelTable
}
