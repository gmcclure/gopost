// Post tests.
package post

import (
	"github.com/gmcclure/gopost/config"
	. "github.com/smartystreets/goconvey/convey"
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
		err = coll.Insert(&Post{"Test Post " + strconv.FormatInt(int64(i), 10), "test_post_" + strconv.FormatInt(int64(i), 10), []byte("This is test post " + strconv.FormatInt(int64(i), 10))})

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

// TestPostGet tests the retrieval of a single post by the post's slug.
func (s *PostSuite) TestPostGet(c *C) {
    Convey("Retrieving a post", c, func() {
        post := Get("test_post_3")
        So(post.Title, ShouldEqual, "Test Post 3")
    })
	// c.Check(post.Title, Equals, "Test Post 3")
}

// TestPostGetAll ensures that all fixtured posts are returned.
func (s *PostSuite) TestPostGetAll(c *C) {
	posts := GetAll()
	c.Check(len(posts), Equals, 11, Commentf("Incorrect number of posts returned"))
}

// TestPostSave simply checks for an error on p.Save(), nothing more.
func (s *PostSuite) TestPostSave(c *C) {
	p := &Post{Title: "Test Post", Body: []byte("This is a test post.")}
	err := p.Save()
	if err != nil {
		c.Errorf("Error saving: %v", err)
	}
}
