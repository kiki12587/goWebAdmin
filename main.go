package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // 引入主题，必须引入，不然报错
	"github.com/gin-gonic/gin"
	"newgoadmin/article"
	"newgoadmin/article_category"
	"newgoadmin/article_comment"
	"newgoadmin/article_label"
	"newgoadmin/article_message"
	"newgoadmin/article_statistics"
	"newgoadmin/article_user_message"
	envconfig "newgoadmin/src/config"
	"newgoadmin/users"
)

type result struct {
	ID       interface{} `json:"id"`
	Text     string      `json:"text"`
	Selected bool        `json:"selected,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
}

func main() {

	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()

	// GoAdmin全局配置，也可以写成一个json，通过json引入
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       envconfig.GetEnv().DatabaseAddr,
				Port:       envconfig.GetEnv().DatabaseProt,
				User:       envconfig.GetEnv().DatabaseUsername,
				Pwd:        envconfig.GetEnv().DatabasePassword,
				Name:       envconfig.GetEnv().DatabaseTablename,
				MaxIdleCon: envconfig.GetEnv().DatabaseMaxIdleCon,
				MaxOpenCon: envconfig.GetEnv().DatabaseMaxOpenCon,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin", // 访问网站的前缀
		// Store 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.CN,
	}

	// 这里引入你需要管理的业务表配置
	// 关于Generators，详见 https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("article", article.GetArticleTable)                                      //文章列表
	adminPlugin.AddGenerator("article_category", article_category.GetArticleCategoryTable)            //文章分类
	adminPlugin.AddGenerator("article_comment", article_comment.GetArticleCommentTable)               //文章评论
	adminPlugin.AddGenerator("article_statistics", article_statistics.GetArticleStatisticsTable)      //文章数据
	adminPlugin.AddGenerator("article_label", article_label.GetArticleLabelTable)                     //文章标签
	adminPlugin.AddGenerator("users", users.GetUsersTable)                                            //作者介绍
	adminPlugin.AddGenerator("article_user_message", article_user_message.GetArticleUserMessageTable) //给作者留言
	adminPlugin.AddGenerator("article_message", article_message.GetArticleMessageTable)               //用户留言

	// 增加配置与插件，使用Use方法挂载到Web框架中
	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	_ = r.Run(":" + envconfig.GetEnv().ServerPort)
}
