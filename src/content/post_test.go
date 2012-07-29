package content

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

func (s *PostSuite) SetUpSuite(c *C) {
	// set up the database for testing
	s.dir = c.MkDir()

	cmdStr := "sqlite3 " + s.dir + "/blog.db < /Users/gmcclure/src/gopost/src/main/blog.sql"

	cmd := exec.Command("sh")
	cmd.Stdin = strings.NewReader(cmdStr)

	if err := cmd.Run(); err != nil {
		c.Fatalf("Error starting command: %v", err)
	}
}

func (s *PostSuite) TestPostSave(c *C) {
	p := &Post{Title: "Test Post", Body: []byte("This is a test post.")}
	if err := p.Save("sqlite3", s.dir+"/blog.db"); err != nil {
		c.Errorf("Error saving: %v", err)
	}
}
