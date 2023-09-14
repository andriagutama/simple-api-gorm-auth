package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-api-gorm-auth/auth"
	"simple-api-gorm-auth/middleware"
	"simple-api-gorm-auth/models"
	"strconv"

	"github.com/gin-gonic/gin" // go get -u github.com/gin-gonic/gin
	_ "github.com/lib/pq"      // go get -u github.com/lib/pq

	"github.com/jinzhu/gorm" // got get -u github.com/jinzhu/gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv" // go get -u github.com/joho/godotenv
)

func postHandler(c *gin.Context, db *gorm.DB) {
	var student models.Student

	c.Bind(&student)
	db.Create(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "success created",
		"data":    student,
	})
}

func getAllHandler(c *gin.Context, db *gorm.DB) {
	var student []models.Student

	db.Find(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    student,
	})
}

func getHandler(c *gin.Context, db *gorm.DB) {
	var student models.Student

	studentId := c.Param("student_id")

	id, _ := strconv.ParseUint(studentId, 10, 64)
	if db.Find(&student, "student_id=?", id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    student,
	})
}

func putHandler(c *gin.Context, db *gorm.DB) {
	var student models.Student

	studentId := c.Param("student_id")

	id, _ := strconv.ParseUint(studentId, 10, 64)
	if db.Find(&student, "student_id=?", id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	var reqStudent models.Student
	c.Bind(&reqStudent)
	db.Model(&student).Update(reqStudent)

	c.JSON(http.StatusOK, gin.H{
		"message": "success update",
		"data":    reqStudent,
	})
}

func deleteHandler(c *gin.Context, db *gorm.DB) {
	studentId := c.Param("student_id")

	var student models.Student
	id, _ := strconv.ParseUint(studentId, 10, 64)
	if db.Find(&student, "student_id=?", id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	db.Delete(&student, "student_id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
	})
}

func setRouter() *gin.Engine {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error env")
	}
	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic("Connection to db failed")
	}

	Migrate(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to simple api gorm auth using golang. course from myskill.id",
		})
	})

	r.POST("login", auth.LoginHandler)

	r.POST("student", middleware.AuthValid, func(ctx *gin.Context) {
		postHandler(ctx, db)
	})

	r.GET("/student", middleware.AuthValid, func(ctx *gin.Context) {
		getAllHandler(ctx, db)
	})

	r.GET("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		getHandler(ctx, db)
	})

	r.PUT("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		putHandler(ctx, db)
	})

	r.DELETE("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		deleteHandler(ctx, db)
	})

	return r
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Student{})

	data := models.Student{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("=================== run seeder user ======================")
		seederUser(db)
	}
}

func seederUser(db *gorm.DB) {
	data := models.Student{
		//Student_id:       1,
		Student_name:     "Dono",
		Student_age:      20,
		Student_address:  "Jakarta",
		Student_phone_no: "0123456789",
	}

	db.Create(&data)
}

func main() {
	r := setRouter()

	r.Run(":8080")
}
