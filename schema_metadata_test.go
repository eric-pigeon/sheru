package sheru

import (
	"github.com/matryer/is"
	"testing"
)

func TestTableName(t *testing.T) {
	is := is.New(t)

	post := Post{}
	changeset := Change(post)

	is.Equal("posts", changeset.metadata.tableName)

	blog := Blog{}
	changeset = Change(blog)
	is.Equal("thoughts", changeset.metadata.tableName)
}

func TestQuotedTableName(t *testing.T) {
	is := is.New(t)

	user := User{}
	changeset := Change(user)

	is.Equal(`"blog"."users"`, changeset.metadata.quotedTableName())

}
