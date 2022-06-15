package controllers


import (
	"github.com/gin-gonic/gin"
)


// PagesController 处理静态页面
type PagesController struct {
}


// Home 首页
func (*PagesController) HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":"枝江",
	})
}


// About 页面
func (*PagesController) AboutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "此博客是记录枝江的，如有反馈或建议，请联系 ava@pg.one",
	})
}


// 404 页面
func (*PagesController) NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{
		"msg": "请求页面未找到:(",
	})
}

