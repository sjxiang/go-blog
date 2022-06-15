package controllers

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"goblog/pkg/database"
)


// 文章相关页面
type ArticlesController struct {
} 


// Show 文章详情页面
func (*ArticlesController) ArticlesShowHandler(c *gin.Context) {

	// 1. 获取 URL 参数
	id := c.Param("id")


	// 2. 读取对应的文章数据
	article, err := database.GetArticleByID(id)


	// 3. 如果出现错误
	if err != nil {

		if err == sql.ErrNoRows {
			
			// 3.1 数据未找到
			c.JSON(404, gin.H{
				"msg": "404 文章未找到",
			})
		} else {

			// 3.2 数据库错误
			if err != nil {
				log.Fatal(err)
			} 
			c.JSON(500, gin.H{
				"msg": "500 服务器内部错误",			
			})
		}

	} else {

		// 4. 读取成功，显示文章
		c.JSON(200, gin.H{
			"msg": "读取成功",
			"data": article,
		})
	}
}
