// Post tests.
package post

import (
	. "launchpad.net/gocheck"
	"os/exec"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PostSuite struct {
	dir string
}

var _ = Suite(&PostSuite{})

// Testing post.go means setting up a datastore.
// SetUpSuite creates a sqlite3 db in a temporary directory for testing.
func (s *PostSuite) SetUpSuite(c *C) {
    // c.Mkdir() creates a directory that will be destroyed after tests.
	// (http://go.pkgdoc.org/launchpad.net/gocheck#C.MkDir)
	s.dir = c.MkDir()

    // Throw the test db in the tmp dir
	cmdStr := "sqlite3 " + s.dir + "/blog.db < /Users/gmcclure/src/gopost/src/gopost/main/gopost.sql"

	cmd := exec.Command("sh")
	cmd.Stdin = strings.NewReader(cmdStr)

	if err := cmd.Run(); err != nil {
		c.Fatalf("Error starting command: %v", err)
	}
}

// TestPostSave simply checks for an error on p.Save(), nothing more.
func (s *PostSuite) TestPostSave(c *C) {
	p := &Post{Title: "Test Post", Body: []byte("This is a test post.")}
	if err := p.Save("sqlite3", s.dir+"/blog.db"); err != nil {
		c.Errorf("Error saving: %v", err)
	}
}
