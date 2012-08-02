// Package post provides gopost's basic post model. The post is conceived as,
// along with a page, a fundmental element of blog content.
package post

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Post is, along with Page, a fundamental unit of content in the blog.
type Post struct {
	Title string
	Body  []byte
}

// Save uses dbDriver and dbName strings to save a post to the database.
func (p *Post) Save(dbDriver, dbName string) error {
	db, err := sql.Open(dbDriver, dbName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	sql := fmt.Sprintf("insert into posts (title, body) values ('%s', '%s')", p.Title, p.Body)
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Printf("%v: %v\n", err, sql)
	}
	return err
}