package content

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Title string
	Body  []byte
}

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
