package database

func (db *databaseManager) ClearDatabase() (err error) {
	_, err = db.dataBase.Exec(`SELECT * FROM func_clear_database()`)
	return
}

func (db *databaseManager) GetDatabase() (database Database, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_get_database()`)
	err = row.Scan(&database.Forum, &database.Post, &database.Thread, &database.User)
	return
}
