package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {

	// get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//connect to database
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s? sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("connect to :", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error ")
	}

	//test ping to database
	if err := db.Ping(); err != nil {
		fmt.Printf("connestion to database failed (DB_HOST: %s: %c\n", dbHost, err)
	} else {
		fmt.Println("connestion to database success", db)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("ouch!")
		return c.HTML(http.StatusOK, "sehat selalu\n")

	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "77"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}

//-----------------------------------------------------------------------------------------------------------------//

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// 	_ "github.com/lib/pq"
// )

// func main() {

// 	// get env
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbName := os.Getenv("DB_NAME")

// 	//connect to database
// 	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s? sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
// 	fmt.Println("connect to :", connStr)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Println("error kyay")
// 	}

// 	//test ping to database
// 	if err := db.Ping(); err != nil {
// 		fmt.Println("connestion to database failed (DB_HOST: %s: %c\n", dbHost, err)
// 	} else {
// 		fmt.Println("connestion to database success", db)
// 	}

// 	e := echo.New()
// 	fmt.Println("someone hit me!")

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.HTML(http.StatusOK, "naon \n")
// 	})

// 	httpPort := os.Getenv("PORT")
// 	if httpPort == "" {
// 		httpPort = "80"
// 	}

// 	e.Logger.Fatal(e.Start(":" + httpPort))

// }
