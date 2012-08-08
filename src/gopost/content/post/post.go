// Package post provides gopost's basic post model. The post is conceived as
// the fundmental unit of blog content.
package post

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopost/config"
)

// Post is the fundamental unit of content in the blog.
type Post struct {
	Title string
	Body  []byte
}

// Cuts down on DB handle boilerplate.
func getDb() *sql.DB {
	db, err := sql.Open(config.DbDriver, config.DbName)
	if err != nil {
		fmt.Println(err)
	}
	// db needs to be closed by caller 
	return db
}

// Provides info associated with SQL queries.
func handleSQLError(err error, sql string) {
	if err != nil {
		fmt.Printf("%v: %v\n", err, sql)
	}
}

// Does the expected, returning all posts in the database. By default
// posts are ordered by post_date.
func (p *Post) ListAll() (*sql.Rows, error) {
	db := getDb()
	defer db.Close()

	sql := fmt.Sprint("select * from posts")
	posts, err := db.Query(sql)
	handleSQLError(err, sql)
	return posts, err
}

// Saves a post to the database using the internal getDB() function and the
// DbDriver and DbName values found in gopost/config.
func (p *Post) Save() error {
	db := getDb()
	defer db.Close()

	sql := fmt.Sprintf("insert into posts (title, body) values ('%s', '%s')", p.Title, p.Body)
	_, err := db.Exec(sql)
	handleSQLError(err, sql)
	return err
}
