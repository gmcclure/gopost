// Post tests.
package post

import (
	"gopost/config"
	"labix.org/v2/mgo"
	. "launchpad.net/gocheck"
	"strconv"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PostSuite struct {
	dir string
}

var _ = Suite(&PostSuite{})

// Setup fixtures -- specifically, 11 test posts.
func (s *PostSuite) SetUpSuite(c *C) {
	config.DbName = config.TestDbName

	session, err := mgo.Dial("localhost")
	if err != nil {
		c.Errorf("Error finding db: %v", err)
	}
	defer session.Close()

	postCount := 11
	coll := session.DB(config.TestDbName).C(config.DbPosts)
	for i := 0; i < postCount; i++ {
		err = coll.Insert(&Post{"Test Post " + strconv.FormatInt(int64(i), 10), []byte("This is test post " + strconv.FormatInt(int64(i), 10))})
		if err != nil {
			c.Errorf("Error inserting test post: %v", err)
		}
	}
}

// Tear down fixtures and any other setup
func (s *PostSuite) TearDownSuite(c *C) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		c.Errorf("Error finding db: %v", err)
	}
	defer session.Close()

	testDb := session.DB(config.TestDbName)
	err = testDb.DropDatabase()
	if err != nil {
		c.Errorf("Error dropping test db: %v", err)
	}
}

// TestPostGetAll ensures that all fixtured posts are returned.
func (s *PostSuite) TestPostGetAll(c *C) {
	posts := GetAll()
	numPosts := len(*posts)
	if numPosts != 11 {
		c.Errorf("Incorrect number of posts returned")
	}
}

// TestPostSave simply checks for an error on p.Save(), nothing more.
func (s *PostSuite) TestPostSave(c *C) {
	p := &Post{Title: "Test Post", Body: []byte("This is a test post.")}
	err := p.Save()
	if err != nil {
		c.Errorf("Error saving: %v", err)
	}
}
