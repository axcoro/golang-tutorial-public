package main
import (
	"database/sql" 
	_ "github.com/go-sql-driver/mysql" 
	"fmt"
)
func main() {
	dbUsername, dbPassword := "dbGoSL_ADMIN", "oMKX98m4Es"
	dbHost, dbName := "10.64.134.51:6612", "dbGoSL"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", connectionString)
	//Check that the database is available and accessible
	err = db.Ping()
	if err == nil {
		fmt.Println( "ok")
	}else {
		fmt.Println( "error: ", err)
	}

	fmt.Println("Probando conexión")
	stmtIns, err := db.Prepare("INSERT INTO company VALUES(?, ?, ?)")
	_,err = stmtIns.Exec(2,"Mercado Libre", "La Punta")
	stmtOut, err := db.Prepare("SELECT name FROM company WHERE id = ? ")
}