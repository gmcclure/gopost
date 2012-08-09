// Package post provides gopost's basic post model. The post is conceived as
// the fundmental unit of blog content.
package post

import (
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"gopost/config"
)

// Post is the fundamental unit of content in the blog.
type Post struct {
	Title string
	Body  []byte
}

// Cuts down on DB handle boilerplate.
func getDb() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil { panic(err) }
	// db needs to be closed by caller 
	return session
}

// Saves a post to the database using the internal getDB() function and the
// DbDriver and DbName values found in gopost/config.
func (p *Post) Save() error {
	session, err := mgo.Dial("localhost")
	if err != nil { panic(err) }
	defer session.Close()

	c := session.DB(config.DbName).C("posts")
	err = c.Insert(&Post{p.Title, p.Body})
	if err != nil { panic(err) }
	return err
}
