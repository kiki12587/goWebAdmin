package article_user_message

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "article_user_message" => http://localhost:9033/admin/info/article_user_message
//
// example end
//
var Generators = map[string]table.Generator{
	"article_user_message": GetArticleUserMessageTable,

	// generators end
}
