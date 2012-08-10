// Package config is a bit of gopost glue, providing a primary means for the
// application's packages to talk to each other while remaining as loosely
// coupled as possible.
package config

var (
	DbName  string = "blog"
	DbPosts string = "posts"
)
