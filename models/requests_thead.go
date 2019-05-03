package models

import (
	"database/sql"
	"time"
)

func (db *dbManager) CreatePost(post Post, created time.Time, id int, forum string) (p Post, err error) {
	row := db.dataBase.QueryRow(`SELECT id, author, thread, forum, message, is_edited, parent, created 
										FROM func_create_post($1::citext, $2::INT, $3::text, $4::INT, $5::citext, $6::TIMESTAMP WITH TIME ZONE)`,
		post.Author, id, post.Message, post.Parent, forum, created)
	err = row.Scan(&p.ID, &p.Author, &p.Thread, &p.Forum,
		&p.Message, &p.IsEdited, &p.Parent, &p.Created)
	return
}

func (db *dbManager) GetThread(slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_get_thread($1::citext, $2::INT)`, slug, threadId)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) UpdateThread(message string, title string, slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_update_thread($1::text, $2::text, $3::citext, $4::INT)`,
		message, title, slug, threadId)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title,
		&thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) CreateOrUpdateVote(vote Vote, slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_create_or_update_vote($1::citext, $2::citext, $3::INT, $4::INT)`,
		vote.Nickname, slug, threadId, vote.Voice)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title,
		&thread.Message, &thread.Votes, &thread.Created)
	return
}


func (db *dbManager) GetPosts(slug string, id int, limit int, since int, sort string, desc bool) (posts []Post, err error) {
	rows, err := getRowsForGetPosts(slug, id, limit, since, sort, desc)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Author, &post.Thread, &post.Forum,
			&post.Message, &post.IsEdited, &post.Parent, &post.Created)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func getRowsForGetPosts(slug string, id int, limit int, since int, sort string, desc bool) (rows *sql.Rows, err error){
	switch sort {
	case "flat":
		rows, err = db.dataBase.Query(
			`SELECT id, author, thread, forum, message, is_edited, parent, created 
					FROM func_get_posts_flat($1::citext, $2::INT, $3::INT, $4::INT, $5::BOOLEAN)`,
			slug, id, limit, since, desc)
		if err != nil {
			return
		}
	case "tree":
		rows, err = db.dataBase.Query(
			`SELECT id, author, thread, forum, message, is_edited, parent, created
					FROM func_get_posts_tree($1::citext, $2::INT, $3::INT, $4::INT, $5::BOOLEAN)`,
			slug, id, limit, since, desc)
		if err != nil {
			return
		}
	case "parent_tree":
		rows, err = db.dataBase.Query(
			`SELECT id, author, thread, forum, message, is_edited, parent, created
					FROM func_get_posts_parent_tree($1::citext, $2::INT, $3::INT, $4::INT, $5::BOOLEAN)`,
			slug, id, limit, since, desc)
		if err != nil {
			return
		}
	default:
		rows, err = db.dataBase.Query(
			`SELECT id, author, thread, forum, message, is_edited, parent, created
					FROM func_get_posts($1::citext, $2::INT, $3::INT, $4::INT, $5::BOOLEAN)`,
			slug, id, limit, since, desc)
		if err != nil {
			return
		}
	}
	return
}