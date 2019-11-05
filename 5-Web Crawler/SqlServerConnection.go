//https://github.com/microsoft/sql-server-samples/blob/master/samples/tutorials/go/crud.go
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

func CreateMasterRecord(db *sql.DB, title string, company string, place string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO dbo.Jobs (Title,Company,Place) VALUES (N'%s',N'%s',N'%s') SELECT SCOPE_IDENTITY();",
		title, company, place)
	var id int64
	err := db.QueryRow(tsql).Scan(&id)

	return int64(id), err
}

func CreateDetailRecord(db *sql.DB, jobId int, key string, value string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO dbo.JobDetails (JobId,[key],[value]) VALUES (%d,N'%s',N'%s');",
		jobId, key, value)
	var id int64
	err := db.QueryRow(tsql).Scan(&id)

	return id, err
}

func GetConnection() *sql.DB {
	condb, errdb := sql.Open("mssql", "server=.;user id=sa;password=123;port=1433;database=JobDb;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	return condb
}
