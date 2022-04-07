package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"golang-sample_api/database"
	"golang-sample_api/routes"
	"golang-sample_api/middleware"
	"gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "userlist"
)

func main() {
	psqlconn := fmt.Sprintf("host= %s port = %d user= %s password = %s dbname= %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	//insertStmt := `insert into "Userlist"("Name", "UserId", "Age", "address", "email", "contact_number") values (' ', 00, 00, ' ', ' ' +63)`
	//_, e := db.Exec(insertStmt)
	//CheckError(e)

	//insertDynStmt := `insert into  "Userlist"("Name", "UserId", "Age", "address", "email", "contact_number") values ($1, $2, $3, $4, $5, $6, $7)`
	//_, e = db.Exec(insertDynStmt, " ", 01, 00, " ", +63)
	//CheckError(e)
}

//func CheckError(err error) {
//	if err != nil {
//		panic(err)
//	}

//}

routers := gin.New()
router.Use(gin.Logger())
routes.userRoutes(router)
router.Use(middleware.Authentication())

routes.FoodRoutes(router)
routes.MenuRoutes(router)
routes.TableRoutes(router)
routes.OrderRoutes(router)
routes.OrderItemRoutes(router)
routes.InvoiceRoutes(router)

router.Run(":" + port)
}
