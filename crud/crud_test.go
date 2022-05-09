package crud

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCrud(t *testing.T) {
	sqlExec()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}
