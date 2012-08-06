// Package post provides gopost's basic post model. The post is conceived as,
// along with a page, a fundmental element of blog content.
package post

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    "gopost/config"
)

// Post is, along with Page, a fundamental unit of content in the blog.
type Post struct {
	Title string
	Body  []byte
}


func getDb() *sql.DB {
	db, err := sql.Open(config.DbDriver, config.DbName)
	if err != nil {
		fmt.Println(err)
	}
	// db needs to be closed by caller 
	return db
}

func handleDbError(err error, sql string) {
	if err != nil {
		fmt.Printf("%v: %v\n", err, sql)
	}
}

// Does the expected, returning all posts in the database. By default
// posts are ordered by date published.
func (p *Post) ListAll() (*sql.Rows, error) {
    db := getDb()
	defer db.Close()

	sql := fmt.Sprint("select * from posts")
    posts, err := db.Query(sql)
    handleDbError(err, sql)
	return posts, err
}

// Saves a post to the database.
func (p *Post) Save() error {
    db := getDb()
	defer db.Close()

	sql := fmt.Sprintf("insert into posts (title, body) values ('%s', '%s')", p.Title, p.Body)
    _, err := db.Exec(sql)
    handleDbError(err, sql)
	return err
}
