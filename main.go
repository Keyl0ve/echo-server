package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"

	"github.com/Keyl0ve/echo-server/routes"

	"github.com/Keyl0ve/echo-server/model"

	"github.com/labstack/echo/v4"
)

// var db *sql.DB

func main() {
	e := echo.New()

	sql, err := InitDatabase("MySQL", 3306, "app", "root", "password")
	if err != nil {
		fmt.Println("error db Open")
		os.Exit(1)
	}
	defer sql.Close()

	// データベース接続情報を直接指定
	// dbUser := "root"
	// dbPassword := "password"
	// dbHost := "MySQL"
	// dbPort := "3306"
	// dbName := "app"

	// connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// var err error
	// db, err = sql.Open("mysql", connString)
	// if err != nil {
	// 	fmt.Printf("Failed to connect to database: %v\n", err)
	// 	return
	// }

	// if err := db.Ping(); err != nil {
	// 	fmt.Printf("Database ping failed: %v\n", err)
	// 	return
	// }

	// Setup API routes
	routes.SetupRoutes(e)

	// Start the server
	e.Start(":8080")
}

func InitDatabase(host string, port uint16, dbname, username, password string) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
	cfg.DBName = dbname
	cfg.User = username
	cfg.Passwd = password
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	model.DB = sql.OpenDB(connector)

	return model.DB, nil
}
