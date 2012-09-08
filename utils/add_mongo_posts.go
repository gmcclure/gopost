package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Post is the fundamental unit of content in the blog.
type Post struct {
	Title string
	Body  []byte
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("posts")
	err = c.Insert(&Post{"Post 1", []byte("This is a very lovely test post.")}, &Post{"Post 2", []byte("This is another very lovely test post.")})
	if err != nil {
		panic(err)
	}

	result := Post{}
	err = c.Find(bson.M{"title": "Post 1"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:", result.Title)
}
