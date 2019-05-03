package database

import "time"

func (db *databaseManager) CreateForum(forum Forum) (f Forum, err error) {

	row := db.dataBase.QueryRow(`SELECT * FROM func_create_forum($1::citext, $2::citext, $3::text)`,
		forum.User, forum.Slug, forum.Title)
	err = row.Scan(&f.IsNew, &f.ID, &f.Slug, &f.User, &f.Title, &f.Posts, &f.Threads)
	return
}

func (db *databaseManager) CreateThread(thread Thread) (t Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM  func_create_thread
 	 ($1::citext, $2::TIMESTAMP WITH TIME ZONE, $3::citext, $4::text, $5::citext, $6::text)`,
		thread.Author, thread.Created, thread.Forum, thread.Message, thread.Slug, thread.Title)
	err = row.Scan(&t.IsNew, &t.ID, &t.Slug, &t.Author, &t.Forum, &t.Title, &t.Message, &t.Votes, &t.Created)
	return
}

func (db *databaseManager) GetForum(slug string) (forum Forum, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_get_forum($1::citext)`, slug)
	err = row.Scan(&forum.IsNew, &forum.ID, &forum.Slug, &forum.User, &forum.Title, &forum.Posts, &forum.Threads)
	return
}

func (db *databaseManager) GetThreads(slug string, since time.Time, desc bool, limit int) (threads []Thread, err error) {

	rows, err := db.dataBase.Query(`SELECT * FROM func_get_threads($1::citext, $2::TIMESTAMP WITH TIME ZONE,
  		$3::BOOLEAN, $4::INT)`, slug, since, desc, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	var thread Thread
	for rows.Next() {
		err = rows.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title,
			&thread.Message, &thread.Votes, &thread.Created)
		if err != nil {
			return
		}
		threads = append(threads, thread)
	}
	return
}

func (db *databaseManager) GetUsers(slug string, since string, desc bool, limit int) (users []User, err error) {

	rows, err := db.dataBase.Query(`SELECT * FROM func_get_users($1::citext, $2::citext, $3::BOOLEAN, $4::INT)`,
		slug, since, desc, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err = rows.Scan(&user.IsNew, &user.ID, &user.Nickname, &user.Email, &user.Fullname, &user.About)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
