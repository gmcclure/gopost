// Post tests.
package post

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PostSuite struct {
	dir string
}

var _ = Suite(&PostSuite{})

// TestPostSave simply checks for an error on p.Save(), nothing more.
func (s *PostSuite) TestPostSave(c *C) {
	p := &Post{Title: "Test Post", Body: []byte("This is a test post.")}
	err := p.Save()
	if err != nil {
		c.Errorf("Error saving: %v", err)
	}
}
