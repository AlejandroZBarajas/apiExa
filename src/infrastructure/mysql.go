package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Connection struct {
	DB *sql.DB
}

var db *sql.DB
var err error

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar archivo .env %v", err)
	}
}

func ConnectDB() {
	loadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	/* fmt.Printf("Usuario: %s, Contraseña: %s, Host: %s, Puerto: %s, Base de Datos: %s\n",
	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"), os.Getenv("DB_NAME")) */

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error al abrir la base de datos %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos %v", err)
	}

	fmt.Println("Conexión exitosa")
	//return &Connection{DB: db}

}

func GetDB() *sql.DB {
	return db
}

func (c *Connection) RunQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return result, nil
}

func (c *Connection) GetDBData(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := c.DB.Query(query, values...)
	if err != nil {
		fmt.Errorf("error: %w", err)
	}
	return rows, nil
}
