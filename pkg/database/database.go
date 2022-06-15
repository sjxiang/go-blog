package database

import (
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB 数据库对象
var DB *sql.DB


func Init() error {

	if err := initDB(); err != nil {
		return err 
	}

	if err := createTable(); err != nil {
		return err
	}

	log.Println(DB.Stats())
	
	return nil
}


func initDB() error {
	
	var err error
	dsn := "root:12345qwe@tcp(127.0.0.1:3306)/blog-service?charset=utf8mb4&parseTime=True&loc=Local"

	// Open() 仅验证参数格式是否正确
	// 初始化全局的 db 对象，连接池对象
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		// log.Fatalf("dsn 格式错误：%v \n", err)
		return err
	}

	// 设置最大连接数
	DB.SetMaxOpenConns(25)  

	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25) 
	
	// 设置每个链接的过期时间 
	DB.SetConnMaxLifetime(time.Minute * 5)  
	

	// 尝试连接（校验 dsn 是否正确）
	if err = DB.Ping(); err != nil {
		// log.Fatalf("连接失败：%v \n", err)
		return err
	}

	return nil
}


func createTable() error {
	
	createArticlesSQL := `
	create table if not exists articles (
		id bigint(20) primary key auto_increment not null,
		title varchar(255) collate utf8mb4_unicode_ci not null,
		body longtext collate utf8mb4_unicode_ci
	)`

	if _, err := DB.Exec(createArticlesSQL); err != nil {
		return err
	} 

	return nil
}


// Article 对应一条文章数据
type Article struct {
	Title, Body string
	ID int64
}

// 辅助函数 1
func GetArticleByID(id string) (Article, error) {
	
	article := Article{}
	querySQL := "select * from articles where id = ?"
	err := DB.QueryRow(querySQL, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}