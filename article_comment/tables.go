package article_comment

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "article_comment" => http://localhost:9033/admin/info/article_comment
//
// example end
//
var Generators = map[string]table.Generator{
	"article_comment": GetArticleCommentTable,

	// generators end
}
