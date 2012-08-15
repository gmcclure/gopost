// Package post provides gopost's basic post model. The post is conceived as
// the fundmental unit of blog content.
package post

import (
	"fmt"
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
	if err != nil {
		panic(err)
	}
	// db needs to be closed by caller 
	return session
}

// Returns all posts in the database.
func GetAll() *[]Post {
	session := getDb()
	defer session.Close()

	var posts []Post

	c := session.DB(config.DbName).C(config.DbPosts)
	iter := c.Find(nil).Iter()
	err := iter.All(&posts)
	if err != nil {
		panic(err)
	}

	return &posts
}

// Saves a post to the database using the internal getDB() function and the
// DbName and DbPosts values found in gopost/config.
func (p *Post) Save() error {
	session := getDb()
	defer session.Close()

	c := session.DB(config.DbName).C(config.DbPosts)
	err := c.Insert(&Post{p.Title, p.Body})
	if err != nil {
		fmt.Println("Error saving post: %v", err)
	}
	return err
}
