package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.Upload)
	}
	pub := r.Group("api/v1")
	{
		pub.POST("user/add", v1.AddUser)
		pub.GET("users", v1.GetUsers)
		pub.GET("categories", v1.GetCategories)
		pub.GET("articles", v1.GetArticles)
		pub.GET("article/info/:id", v1.GetArticleInfo)
		pub.GET("article/list/:id", v1.GetCateArticle)
		pub.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
