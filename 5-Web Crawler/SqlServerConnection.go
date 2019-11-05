package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestConnection() {
	condb := GetConnection()

	var (
		Id   int
		Name string
	)
	rows, err := condb.Query("select * from dbo.TestTable")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&Id, &Name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(Id)
		log.Println(Name)
	}

	defer condb.Close()
}

func CreateRecord(db *sql.DB, title string, company string, place string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO dbo.Jobs (Title,Company,Place) VALUES (N'%s',N'%s',N'%s');",
		title, company, place)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func GetConnection() *sql.DB {
	condb, errdb := sql.Open("mssql", "server=.;user id=sa;password=hN1234!@#$;port=1433;database=JobDb;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	return condb
}
