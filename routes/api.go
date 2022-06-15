package routes

import (
	"github.com/gin-gonic/gin"

	"goblog/app/http/controllers"
)

// 注册 api 路由
func RegisterApiRoutes(r *gin.Engine) {
	
	// ===
	pc := new(controllers.PagesController)

	// 首页
	r.GET("/", pc.HomeHandler) 
	
	// 相关
	r.GET("/about", pc.AboutHandler) 
	

	// ====
	ac := new(controllers.ArticlesController)

	// 文章详情
	r.GET("/articles/:id", ac.ArticlesShowHandler)

	// // 文章列表
	// r.GET("/articles", articlesIndexHandler)

	// // 创建文章
	// r.GET("/articles/create", articlesCreateHandler)

	// // 保存文章
	// r.POST("/articles", articlesStoreHandler)

	// // 删除文章
	// r.POST("/articles/:id/delete", articlesDeleteHandler)

	// // 测试中间件
	// r.GET("/check", checkHandler)

	// 自定义 404  
	r.NoRoute(pc.NotFoundHandler)
}