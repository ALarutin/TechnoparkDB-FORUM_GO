package database

func (db *databaseManager) GetUser(nickname string) (user User, err error) {

	row := db.dataBase.QueryRow(
		`SELECT * FROM func_get_user($1::citext)`,
		nickname)
	err = row.Scan(&user.IsNew, &user.ID, &user.Nickname, &user.Email, &user.Fullname, &user.About)
	return
}

func (db *databaseManager) CreateUser(user User) (users []User, err error) {
	rows, err := db.dataBase.Query(`SELECT * FROM func_create_user($1::citext, $2::citext, $3::text, $4::text)`,
		user.Nickname, user.Email, user.Fullname, user.About)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.IsNew, &user.ID, &user.Nickname, &user.Email, &user.Fullname, &user.About)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (db *databaseManager) UpdateUser(user User) (u User, err error) {
	row := db.dataBase.QueryRow(
		`SELECT * FROM func_update_user($1::citext, $2::citext, $3::text, $4::text)`,
		user.Nickname, user.Email, user.Fullname, user.About)
	err = row.Scan(&u.IsNew, &u.ID, &u.Nickname, &u.Email, &u.Fullname, &u.About)
	return
}
