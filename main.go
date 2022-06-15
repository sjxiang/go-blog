package main

import (
	// "log"
	"os"
	// "strconv"

	// "unicode/utf8"


	// "github.com/gin-gonic/gin"

	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"


	"goblog/pkg/server"
	"goblog/pkg/database"
)



func main() {

	if err := database.Init(); err != nil {
		os.Exit(-1)
	}

	// db = database.DB


	r := server.NewRouter()
	r.Run(":3001")

}


// func articlesIndexHandler(c *gin.Context) {

// 	// 1. 执行查询语句，返回一个结果集
// 	rows, err := db.Query("select * from articles")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()


// 	// 2. 循环读取结果
// 	var articles []Article

// 	for rows.Next() {
		
// 		var article Article

// 		// 2.1 扫描每一行的结果并赋值到一个 article 对象中
// 		err := rows.Scan(&article.ID, &article.Title, &article.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// 2.2 将 article 追加到 articles 的这个数组中
// 		articles = append(articles, article)
// 	}

// 	// 2.3 检测遍历时，是否发生错误
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}


// 	// 3. 加载

// 	c.JSON(200, gin.H{
// 		"msg": "hello，欢迎来到 go-blog",
// 		"blog-list": articles,
// 	})
// }

// // Article 对应一条文章数据
// type Article struct {
// 	Title, Body string
// 	ID int64
// }



// func articlesCreateHandler(c *gin.Context) {

// 	c.JSON(200, gin.H{
// 		"msg": "这里写博客", 
// 	})
// }


// // 创建 blog 表单数据
// type ArticlesFormData struct {
// 	Title, Body 	string
// 	Errors 			map[string]string
// } 


// func articlesStoreHandler(c *gin.Context) {

// 	title := c.PostForm("title")
// 	body := c.PostForm("body")

// 	errors := make(map[string]string)

// 	// 验证标题
// 	if title == "" {
// 		errors["title"] = "标题不能为空"
//  	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
// 		errors["title"] = "标题长度需介于 3 - 40 字符"
// 	}


// 	// 验证内容
// 	if body == "" {
// 		errors["body"] = "内容不能为空"
//  	} else if utf8.RuneCountInString(body) < 10 {
// 		errors["body"] = "内容长度需大于 10 个字符"
// 	}


// 	data := ArticlesFormData{
// 		Title: title,
// 		Body: body,
// 		Errors: errors,
// 	}

// 	// 检查是否有错误发生
// 	if len(errors) == 0 {
		
// 		lastInsertID, err := saveArticleToDB(title, body)
		
// 		if lastInsertID > 0 {
// 			c.JSON(200, gin.H{
// 				"msg": "插入成功，ID 为 "+strconv.FormatInt(lastInsertID, 10),
// 			})
		
// 		} else {

// 			if err != nil {
// 				log.Fatal(err)
// 			}
		
// 			c.JSON(500, gin.H{
// 				"msg": "500 服务器内部错误",
// 			})
// 		}

// 	} else {
		
// 		// 失败，展示错误原因与详情
// 		c.JSON(200, data)
// 	}
// }




// func saveArticleToDB(title string, body string) (int64, error) {

// 	// 变量初始化
// 	var (
// 		id int64
// 		err error
// 		rs sql.Result
// 		stmt *sql.Stmt
// 	)

// 	// 1. 获取一个 prepare 声明语句
// 	stmt, err = db.Prepare("insert into articles (title, body) values (?, ?)")
// 	if err != nil {
// 		return 0, err
// 	}

// 	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 链接
// 	defer stmt.Close()

// 	// 3. 执行请求，传参进入绑定的内容
// 	rs, err = stmt.Exec(title, body)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// 4. 插入成功的话，会返回自增 ID
// 	if id, err = rs.LastInsertId(); id > 0 {
// 		return id, nil
// 	}

// 	return 0, err
// }


// func articlesDeleteHandler(c *gin.Context) {
	
// 	// 1. 获取 URL 参数
// 	id := c.Param("id")


// 	// 2. 读取对应的文章数据
// 	article, err := getArticleByID(id)


// 	// 3. 如果出现错误
// 	if err != nil {

// 		if err == sql.ErrNoRows {
			
// 			// 3.1 数据未找到
// 			c.JSON(404, gin.H{
// 				"msg": "404 文章未找到",
// 			})
	
// 		} else {

// 			// 3.2 数据库错误
// 			log.Fatal(err)

// 			c.JSON(500, gin.H{
// 				"msg": "500 服务器内部错误",			
// 			})
// 		}

// 	} else {

// 		// 4. 未出现错误，执行删除操作
// 		rowsAffected, err := article.Delete()
		
// 		// 4.1 发生错误
// 		if err != nil {
			
// 			log.Fatal(err)
// 			// c.JSON(500, gin.H{
// 			// 	"msg": "500 服务器内部错误",
// 			// })
		
// 		} else {

// 			// 4.2 未发生错误
// 			if rowsAffected > 0 {

// 				// 重定向到文章列表页
// 				c.JSON(200, gin.H{
// 					"msg": "删除成功",
// 				})

// 			} else {

// 				// edge case
// 				c.JSON(404, gin.H{
// 					"msg": "404 nmsl 文章未找到",
// 				})
				
// 			}

// 		}

// 	}
// }

// // 从数据库中删除单条记录
// func (a Article) Delete() (rowsAffected int64, err error) {
	
// 	deleteSQL := "delete from articles where id = " + strconv.FormatInt(a.ID, 10)
	
// 	rs, err := db.Exec(deleteSQL)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// 删除成功，跳转文章列表页
// 	if n, _ := rs.RowsAffected(); n > 0 {
// 		return n, nil
// 	}

// 	return 0, nil
// }


// func checkHandler(c *gin.Context) {
// 	ava := c.MustGet("ava")
// 	log.Println(ava)
// }


