package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/fixtures"
)

func TestTreeEntriesRelation(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	r, err := git.NewFilesystemRepository(f.DotGit().Base())
	assert.Nil(err)

	db := NewDatabase("foo", r)
	assert.NotNil(db)

	relations := db.Relations()
	rel, ok := relations[treeEntriesRelationName]
	assert.True(ok)
	assert.NotNil(rel)
	assert.Equal(treeEntriesRelationName, rel.Name())
	assert.Equal(0, len(rel.Children()))

	iter, err := rel.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	row, err := iter.Next()
	assert.Nil(err)
	assert.NotNil(row)

	fields := row.Fields()
	assert.NotNil(fields)
	assert.Len(fields, 4)
	assert.IsType("", fields[0])
	assert.IsType("", fields[1])
	assert.IsType("", fields[2])
	assert.IsType("", fields[3])
}
