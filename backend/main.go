package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 1. 定义与数据库对应的结构体
type Student struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
	Contact   string `json:"contact"`
}

func initDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	// 自动迁移表结构（如果 init.sql 没生效，这句也能帮你建表）
	DB.AutoMigrate(&Student{})
}

func main() {
	initDB()
	r := gin.Default()

	// 2. 配置 CORS（解决跨域问题）
	r.Use(cors.Default())

	// 3. 编写接收数据的 POST 接口
	r.POST("/api/submit", func(c *gin.Context) {
		var student Student
		// 将前端传来的 JSON 数据绑定到 student 结构体上
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "数据格式不正确"})
			return
		}

		// 存入数据库
		if err := DB.Create(&student).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败，可能是学号重复"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "提交成功！", "data": student})
	})

	r.Run(":3000")
}