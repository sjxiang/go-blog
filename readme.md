### go - blog


```
模仿 learnku.com 上的 Go1 Go 实战：web 入门

    原来
        Gorilla Mux
        template + HTML 
        
        
    现在
        Gin
        Json
        
``` 


======

中间件

    客户端
       |  
      \|/
    
    中间件 A 
    
       |
      \|/
    
    业务逻辑处理



======

MySQL

添加 MySQL 驱动
$ go get github.com/go-sql-driver/mysql


dsn := "root:12345qwe@tcp(127.0.0.1:3306)/blog-service?charset=utf8mb4&parseTime=True&loc=Local"


$ mysql -u root -p 
(root@localhost) [(none)]> create database blog-service character set utf8mb4 collate utf8mb4_unicode_ci;

(root@localhost) [blog-service]> create table if not exists articles (
		id bigint(20) primary key auto_increment not null,
		title varchar(255) collate utf8mb4_unicode_ci not null,
		body longtext collate utf8mb4_unicode_ci
	);


sql.Result 

type Result interface {
 
    // 使用 INSERT 向数据插入记录，数据表有自增 ID 时，该函数有返回值
    LastInsertId() (int64, error)
 
    // 表示影响的数据表行数，常用于 UPDATE/DELETE 等 SQL 语句中
    RowsAffected() (int64, error)

}



增
    insert into

删

改
    form 提交内容

    根据 id 查出那条记录
        有
            提取表单参数
            验证表单
            

    验证数据

    更新

查



netstat -nlp | grep :3000 | awk '{print $7}'
首页
    文章列表

编辑博客页面


         
======

MVC

    1. 浏览器输入 URL，回车
    2. 服务端 router 将 URL 请求映射到指定控制器 handler 上
    3. 控制器 handler 收到请求，开始进行处理；
    如果视图 view 需要动态数据组装渲染，则控制器 handler 会从模型 model 中读取数据

    4. 数据读取完毕，数据发送给视图 view，组装渲染（c.JSON(200, gin.H{})）。
 

不同业务场景
    
    不同控制器

        PageController
        ArticlesController

======

测试 

$ go get github.com/stretchr/testify

$ go test ./tests -v



======

gorm

$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
