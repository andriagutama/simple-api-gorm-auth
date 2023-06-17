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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "just4andria"
	dbname   = "db_reff"
)

func postHandler(c *gin.Context, db *gorm.DB) {
	var student models.Student

	/*if c.Bind(&student) == nil {
		_, err := db.Exec("insert into students (student_name, student_age, student_address, student_phone_no) values ($1, $2, $3, $4)",
			student.Student_name,
			student.Student_age,
			student.Student_address,
			student.Student_phone_no,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success created",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "error",
	})*/

	c.Bind(&student)
	db.Create(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "success created",
		"data":    student,
	})
}

/*func rowToStruct(rows *sql.Rows, dest interface{}) error {
	destv := reflect.ValueOf(dest).Elem()
	args := make([]interface{}, destv.Type().Elem().NumField())

	for rows.Next() {
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}*/

func getAllHandler(c *gin.Context, db *gorm.DB) {
	var student []models.Student

	/*row, err := db.Query("SELECT * FROM students ORDER BY student_id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	rowToStruct(row, &student)

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    student,
	})*/

	db.Find(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    student,
	})
}

func getHandler(c *gin.Context, db *gorm.DB) {
	var student models.Student

	studentId := c.Param("student_id")

	/*row, err := db.Query("SELECT * FROM students WHERE student_id = $1",
		studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	rowToStruct(row, &student)

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    student,
	})*/

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

	/*if c.Bind(&student) == nil {
		_, err := db.Exec("update students set student_name=$1 where student_id=$2",
			student.Student_name, studentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success update",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "error",
	})*/

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

	/*_, err := db.Exec("delete from students where student_id = $1",
		studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
	})*/

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
	//conn := "postgresql://postgres:just4andria@127.0.0.1/db_tutorial?sslmode=disable"
	//conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

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

	r.POST("student", func(ctx *gin.Context) {
		postHandler(ctx, db)
	})

	r.GET("/student", middleware.AuthValid, func(ctx *gin.Context) {
		getAllHandler(ctx, db)
	})

	r.GET("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		getHandler(ctx, db)
	})

	r.PUT("/student/:student_id", func(ctx *gin.Context) {
		putHandler(ctx, db)
	})

	r.DELETE("/student/:student_id", func(ctx *gin.Context) {
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

	r.Run()
}
