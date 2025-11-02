package main

import (
	app "practise/applications"
)

func main() {
	db := app.DbService()
	db.ConnectSqlite()
	defer db.CloseSqlite()
	db.Run()
}
