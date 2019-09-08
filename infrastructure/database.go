package infrastructure
import (
	"log"
	//initialize sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

//InitDB ..
func InitDB(driver, connString string) *sqlx.DB {
	
	db, err := sqlx.Connect(driver, connString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}