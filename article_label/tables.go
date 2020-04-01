package article_label

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "article_label" => http://localhost:9033/admin/info/article_label
//
// example end
//
var Generators = map[string]table.Generator{
	"article_label": GetArticleLabelTable,

	// generators end
}
